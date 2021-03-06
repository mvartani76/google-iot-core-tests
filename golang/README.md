# Google Cloud IoT Core GoLang MQTT example

# Setup Code Environment (32 bit Raspberry Pi)

## Look up the most recent go version:
https://golang.org/dl/

  ``` ctrl+f armv6 ```

Latest as of now is 1.13.5

## Copy the download link:
https://dl.google.com/go/go1.10.2.linux-armv6l.tar.gz

## Paste the following, line by line:
```
  wget https://dl.google.com/go/go1.13.5.linux-armv6l.tar.gz

  sudo tar -C /usr/local -xvf go1.13.5.linux-armv6l.tar.gz

  cat >> ~/.bashrc << 'EOF'

  export GOPATH=$HOME/go

  export PATH=/usr/local/go/bin:$PATH:$GOPATH/bin

  EOF

  source ~/.bashrc
```

## Verify Working Install by checking version
  ``` go version ```
  
  ![Golang version](https://github.com/mvartani76/google-iot-core-tests/blob/master/images/golang-version-test-output.png "Nodejs Golang Version - Successful Install")

# Setup Device

## Generate a device key pair (on device)
Open a terminal window and run the following multi-line command to create an RS256 key:

    openssl req -x509 -newkey rsa:2048 -keyout rsa_private.pem -nodes \
    -out rsa_cert.pem -subj "/CN=unused"

In the following section, you'll add a device to the registry and associate the public key with the device.

## Install Requirements

The MQTT and jwt vars are not currently available so we need to get the following packages as shown below:

```
# jwt
go get github.com/dgrijalva/jwt-go
# MQTT
go get github.com/eclipse/paho.mqtt.golang
```

## Add a device to the registry (console)
1. On the **Registries** page, select ```my-registry```.

2. Select the **Devices** tab and click **Create a device**.

3. Enter ```golang-device``` for the **Device ID**.

4. Select **Allow** for **Device communication**.

5. Add the public key information to the **Authentication** fields.
    - Copy the contents of rsa_cert.pem to the clipboard. Make sure to include the lines that say -----BEGIN CERTIFICATE----- and -----END CERTIFICATE-----.
    - Select **RS256_X509** for the **Public key format**.
    - Paste the public key in the **Public key value** box.

6. The **Device metadata** field is optional; leave it blank.

7. Click **Create**.

You've just added a device to your registry. The RS256_X509 key appears on the Device details page for your device.

# Running the sample (on device)

The following command summarizes the sample usage:

```
  go run main.go \
  --device=golang-device \
  --project=<your-project-id> \
  --registry=my-registry \
  --region=us-central1 \
  --ca_certs=./roots.pem \
  --private_key=./rsa_private.pem
```

# Verification of Running Code

If everything is setup/configured correctly, the following will show what the outputs will look like on the device.

![Golang Main Output](https://github.com/mvartani76/google-iot-core-tests/blob/master/images/golang-main-working-output.png "Golang Main Output")

# Useful Links
https://medium.com/google-cloud/google-cloud-iot-core-golang-b130f65951ba
https://github.com/GoogleCloudPlatform/golang-samples/tree/master/iot
https://github.com/dgrijalva/jwt-go
