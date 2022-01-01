<script lang="ts" setup>
import * as tus from "tus-js-client";
import { ref } from "vue";
const percent = ref("");

const onInputChange = (e: Event) => {
  // Get the selected file from the input element
  let target = e.target as HTMLInputElement;

  if (!target || !target.files) return;

  let file = target.files[0];

  // Create a new tus uploads
  let upload = new tus.Upload(file, {
    endpoint: "/files/",
    retryDelays: [0, 3000, 5000, 10000, 20000],
    metadata: {
      filename: file.name,
      filetype: file.type,
    },
    onError: function (error) {
      console.log("Failed because: " + error);
    },
    onProgress: function (bytesUploaded, bytesTotal) {
      let percentage = ((bytesUploaded / bytesTotal) * 100).toFixed(2);
      console.log(bytesUploaded, bytesTotal, percentage + "%");
      percent.value = percentage;
    },
    onSuccess: function() {
      console.log("success");
      console.log(
        "Download %s from %s",
        upload.file instanceof File ? upload.file.name : "",
        upload.url
      );
    },
    onAfterResponse: function (req, res) {
      let url = req.getURL()
      let value = res.getHeader("X-My-Header")
      console.log(upload.file)
      console.log(res)
      console.log(`Request for ${url} responded with ${value}`)
    },
  });

  // Check if there are any previous uploads to continue.
  // uploads.findPreviousUploads().then(function (previousUploads) {
  //   // Found previous uploads so we select the first one.
  //   if (previousUploads.length) {
  //     uploads.resumeFromPreviousUpload(previousUploads[0]);
  //   }
  //
  //   // Start the uploads
  //   uploads.start();
  // });
    upload.start();
};
</script>

<template>
  <div>
    <input type="file" @change="onInputChange" />
    <span>{{ percent }}%</span>
  </div>
</template>
