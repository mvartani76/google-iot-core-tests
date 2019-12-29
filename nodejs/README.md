<img src="https://avatars2.githubusercontent.com/u/2810941?v=3&s=96" alt="Google Cloud Platform logo" title="Google Cloud Platform" align="right" height="96" width="96"/>

# Google Cloud IoT Core NodeJS MQTT example

This sample app publishes data to Cloud Pub/Sub using the MQTT bridge provided
as part of Google Cloud IoT Core.

Note that before you can run this sample, you must register a device as
described in the parent README. For the gateway samples, you must register and bind
a device as described in the [Cloud IoT gateway docs](https://cloud.google.com/iot/docs/how-tos/gateways/#setup).

# Setup Code Environment

Run the following command to install nvm (bash script for managing installations of Node.js and npm):

    wget -qO- https://raw.githubusercontent.com/nvm-sh/nvm/v0.35.2/install.sh | bash

You will probably need to close and restart the terminal to have changes take effect.

Run the following command to install the latest version of Node.js:

    nvm install stable

Run the following command to install the library dependencies for NodeJS:

    npm install

# Configure Cloud Console Registry and Device

## Create a Device Registry

1. Go to the [Google Cloud IoT Core page in Cloud Console](https://console.cloud.google.com/iot).

2. Click **Create registry.**

3. Enter ```my-registry``` for the **Registry ID.**

4. If you're in the US, select **us-central1** for the **Region.** If you're outside the US, select your preferred region.

5. Select **MQTT** for the **Protocol**.

6. In the **Default telemetry topic** dropdown list, select **Create a topic.**

7. In the **Create a topic** dialog, enter ```my-device-events``` in the **Name** field.

8. Click **Create** in the **Create a topic** dialog.

9. The **Device state topic** and **Certificate value** fields are optional, so leave them blank.

10. Click **Create** on the Cloud IoT Core page.

You've just created a device registry with a Cloud Pub/Sub topic for publishing device telemetry events.

## Generate a device key pair (on device)

Open a terminal window and run the following multi-line command to create an RS256 key:

```
    openssl req -x509 -newkey rsa:2048 -keyout rsa_private.pem -nodes \
    -out rsa_cert.pem -subj "/CN=unused"
```

In the following section, you'll add a device to the registry and associate the public key with the device.

## Add a device to the registry
1. On the **Registries** page, select ```my-registry```.

2. Select the **Devices** tab and click **Create a device**.

3. Enter ```my-device``` for the **Device ID**.

4. Select **Allow** for **Device communication**.

5. Add the public key information to the **Authentication** fields.
    - Copy the contents of rsa_cert.pem to the clipboard. Make sure to include the lines that say -----BEGIN CERTIFICATE----- and -----END CERTIFICATE-----.
    - Select **RS256_X509** for the **Public key format**.
    - Paste the public key in the **Public key value** box.

6. The **Device metadata** field is optional; leave it blank.

7. Click **Create**.

You've just added a device to your registry. The RS256_X509 key appears on the Device details page for your device.

# Running the sample

The following command summarizes the sample usage:

    Usage: cloudiot_mqtt_example_nodejs [command] [options]

    Commands:
        mqttDeviceDemo              Example Google Cloud IoT Core MQTT device connection demo.
        sendDataFromBoundDevice     Demonstrates sending data from a gateway on behalf of a bound device.
        listenForConfigMessages     Demonstrates listening for config messages on a gateway client of a bound device.
        listenForErrorMessages      Demonstrates listening for error messages on a gateway.

    Options:

        --projectId           The Project ID to use. Defaults to the value of the GCLOUD_PROJECT or GOOGLE_CLOUD_PROJECT
                            environment variables.
        --cloudRegion         GCP cloud region.
        --registryId          Cloud IoT registry ID.
        --deviceId            Cloud IoT device ID.
        --privateKeyFile      Path to private key file.
        --algorithm           Encryption algorithm to generate the JWT.
        --numMessages         Number of messages to publish.
        --tokenExpMins        Minutes to JWT token expiration.
        --mqttBridgeHostname  MQTT bridge hostname.
        --mqttBridgePort      MQTT bridge port.
        --messageType         Message type to publish.
        --help                Show help

For example, if your project ID is `blue-jet-123`, your service account
credentials are stored in your home folder in creds.json and you have generated
your credentials using the shell script provided in the parent folder, you can
run the following examples:

    node cloudiot_mqtt_example_nodejs.js mqttDeviceDemo \
        --projectId=blue-jet-123 \
        --cloudRegion=us-central1 \
        --registryId=my-registry \
        --deviceId=my-device \
        --privateKeyFile=../rsa_private.pem \
        --algorithm=RS256

    node cloudiot_mqtt_example_nodejs.js sendDataFromBoundDevice \
        --projectId=blue-jet-123 \
        --cloudRegion=us-central1 \
        --registryId=my-registry \
        --gatewayId=my-gateway \
        --deviceId=my-device \
        --privateKeyFile=../rsa_private.pem \
        --algorithm=RS256

    node cloudiot_mqtt_example_nodejs.js listenForConfigMessages \
        --projectId=blue-jet-123 \
        --cloudRegion=us-central1 \
        --registryId=my-registry \
        --gatewayid=my-gateway \
        --deviceId=my-device \
        --privateKeyFile=../rsa_private.pem \
        --algorithm=RS256
        --clientDuration=60000

# Sending a configuration update

For `listenForConfigMessages` example, try sending a config update to the device while the client is running. This can be done via the Google Cloud IoT Core UI or through the command line with the following command.

    gcloud iot devices configs update --region=us-central1 --registry=my-registry --device=my-device --config-data="testing"

# Reading the messages written by the sample client

1. Create a subscription to your topic.

        gcloud pubsub subscriptions create \
            projects/your-project-id/subscriptions/my-subscription \
            --topic device-events

2. Read messages published to the topic

        gcloud pubsub subscriptions pull --auto-ack \
            projects/my-iot-project/subscriptions/my-subscription

# Verification of Running Code

If everything is setup/configured correctly, the following will show what the outputs will look like on the device.

## MQTTDeviceDemo

![Nodejs MQTTDeviceDemo Output](https://github.com/mvartani76/google-iot-core-tests/blob/master/images/nodejs-mqttDeviceDemo-working-output.png "Nodejs MQTTDeviceDemo Output")
