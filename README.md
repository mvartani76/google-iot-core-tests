# google-iot-core-tests
Repository of code for testing Google IoT Core using various languages

## Set up your local environment and install prerequisites

[Install and initialize the Cloud SDK.](https://cloud.google.com/sdk/docs/) Cloud IoT Core requires version 173.0.0 or higher of the SDK.

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
