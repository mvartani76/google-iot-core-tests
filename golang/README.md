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
