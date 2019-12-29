package main

import (
	"crypto/tls"
	"crypto/x509"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	MQTT "github.com/eclipse/paho.mqtt.golang"
)

var (
	deviceID = flag.String("device", "", "Cloud IoT Core Device ID")
	bridge   = struct {
		host *string
		port *string
	}{
		flag.String("mqtt_host", "mqtt.googleapis.com", "MQTT Bridge Host"),
		flag.String("mqtt_port", "8883", "MQTT Bridge Port"),
	}
	projectID  = flag.String("project", "", "GCP Project ID")
	registryID = flag.String("registry", "", "Cloud IoT Registry ID (short form)")
	region     = flag.String("region", "", "GCP Region")
	certsCA    = flag.String("ca_certs", "", "Download https://pki.google.com/roots.pem")
	privateKey = flag.String("private_key", "", "Path to private key file")
)

func main() {
	log.Println("[main] Entered")

	log.Println("[main] Flags")
	flag.Parse()

	log.Println("[main] Loading Google's roots")
	certpool := x509.NewCertPool()
	pemCerts, err := ioutil.ReadFile(*certsCA)
	if err == nil {
		certpool.AppendCertsFromPEM(pemCerts)
	}

	log.Println("[main] Creating TLS Config")

	config := &tls.Config{
		RootCAs:            certpool,
		ClientAuth:         tls.NoClientCert,
		ClientCAs:          nil,
		InsecureSkipVerify: true,
		Certificates:       []tls.Certificate{},
		MinVersion:         tls.VersionTLS12,
	}

	clientID := fmt.Sprintf("projects/%v/locations/%v/registries/%v/devices/%v",
		*projectID,
		*region,
		*registryID,
		*deviceID,
	)

	log.Println("[main] Creating MQTT Client Options")
	opts := MQTT.NewClientOptions()

	broker := fmt.Sprintf("ssl://%v:%v", *bridge.host, *bridge.port)
	log.Printf("[main] Broker '%v'", broker)

	opts.AddBroker(broker)
	opts.SetClientID(clientID).SetTLSConfig(config)

	opts.SetUsername("unused")

	token := jwt.New(jwt.SigningMethodRS256)
	token.Claims = jwt.StandardClaims{
		Audience:  *projectID,
		IssuedAt:  time.Now().Unix(),
		ExpiresAt: time.Now().Add(24 * time.Hour).Unix(),
	}

	log.Println("[main] Load Private Key")
	keyBytes, err := ioutil.ReadFile(*privateKey)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("[main] Parse Private Key")
	key, err := jwt.ParseRSAPrivateKeyFromPEM(keyBytes)
	if err != nil {
		log.Fatal(err)
	}

	log.Println("[main] Sign String")
	tokenString, err := token.SignedString(key)
	if err != nil {
		log.Fatal(err)
	}

	opts.SetPassword(tokenString)

	// Incoming
	opts.SetDefaultPublishHandler(func(client MQTT.Client, msg MQTT.Message) {
		fmt.Printf("[handler] Topic: %v\n", msg.Topic())
		fmt.Printf("[handler] Payload: %v\n", msg.Payload())
	})

	log.Println("[main] MQTT Client Connecting")
	client := MQTT.NewClient(opts)
	if token := client.Connect(); token.Wait() && token.Error() != nil {
		log.Fatal(token.Error())
	}

	topic := struct {
		config    string
		telemetry string
	}{
		config:    fmt.Sprintf("/devices/%v/config", *deviceID),
		telemetry: fmt.Sprintf("/devices/%v/events", *deviceID),
	}

	log.Println("[main] Creating Subscription")
	client.Subscribe(topic.config, 0, nil)

	log.Println("[main] Publishing Messages")
	for i := 0; i < 10; i++ {
		log.Printf("[main] Publishing Message #%d", i)
		token := client.Publish(
			topic.telemetry,
			0,
			false,
			fmt.Sprintf("Message: %d", i))
		token.WaitTimeout(5 * time.Second)
	}

	log.Println("[main] MQTT Client Disconnecting")
	client.Disconnect(250)

	log.Println("[main] Done")
}