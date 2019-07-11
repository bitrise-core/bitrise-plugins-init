package reactnative

const deployWorkflowDescription = `## Configure Android part of the deploy workflow

To generate a signed APK:

1. Open the **Workflow** tab of your project on Bitrise.io
1. Add **Sign APK step right after Android Build step**
1. Click on **Code Signing** tab
1. Find the **ANDROID KEYSTORE FILE** section
1. Click or drop your file on the upload file field
1. Fill the displayed 3 input fields:
1. **Keystore password**
1. **Keystore alias**
1. **Private key password**
1. Click on **[Save metadata]** button

That's it! From now on, **Sign APK** step will receive your uploaded files.

## Configure iOS part of the deploy workflow

To generate IPA:

1. Open the **Workflow** tab of your project on Bitrise.io
1. Click on **Code Signing** tab
1. Find the **PROVISIONING PROFILE** section
1. Click or drop your file on the upload file field
1. Find the **CODE SIGNING IDENTITY** section
1. Click or drop your file on the upload file field
1. Click on **Workflows** tab
1. Select deploy workflow
1. Select **Xcode Archive & Export for iOS** step
1. Open **Force Build Settings** input group
1. Specify codesign settings
Set **Force code signing with Development Team**, **Force code signing with Code Signing Identity**  
and **Force code signing with Provisioning Profile** inputs regarding to the uploaded codesigning files
1. Specify manual codesign style
If the codesigning files, are generated manually on the Apple Developer Portal,  
you need to explicitly specify to use manual coedsign settings  
(as ejected rn projects have xcode managed codesigning turned on).  
To do so, add 'CODE_SIGN_STYLE="Manual"' to 'Additional options for xcodebuild call' input

## To run this workflow

If you want to run this workflow manually:

1. Open the app's build list page
2. Click on **[Start/Schedule a Build]** button
3. Select **deploy** in **Workflow** dropdown input
4. Click **[Start Build]** button

Or if you need this workflow to be started by a GIT event:

1. Click on **Triggers** tab
2. Setup your desired event (push/tag/pull) and select **deploy** workflow
3. Click on **[Done]** and then **[Save]** buttons

The next change in your repository that matches any of your trigger map event will start **deploy** workflow.
`
