<template>
  <div class="relative flex flex-col gap-4">
    <p class="text-gold-light">Upload a PDF of your session notes.</p>
    <input-wrapper
      v-model="file"
      :validation-rules="{
        accept: '.pdf,application/pdf',
        required: true,
        maxSize: 10 * 1024 * 1024
      }"
      type="file"
      @input="fileError = null"
      @validation-error="fileError = $event"
    />
    <p class="text-red-500 text-sm">{{ fileError }}</p>

    <div class="w-full flex justify-end items-center gap-2 mt-4">
      <button
        :disabled="fileError !== null || file === null"
        class="px-4 py-2 transition-all bg-gold hover:bg-gold-light rounded text-ink disabled:bg-gray-400"
        @click="$emit('upload', file)"
      >
        Upload
      </button>
      <button
        class="px-4 py-2 transition-all bg-ember hover:bg-ember-light rounded text-gold-light"
        @click="$emit('close')"
      >
        Cancel
      </button>
    </div>
  </div>
</template>

<script>
  import InputWrapper from "@/components/forms/InputWrapper.vue";

  export default {
    name: "UploadSessionNotesModal",
    components: {InputWrapper},
    data() {
      return {
        file: null,
        fileError: null,
      }
    },
  }
</script>

<style scoped>

</style>