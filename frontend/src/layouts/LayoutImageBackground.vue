<template>
  <div class="relative min-h-screen min-w-screen overflow-x-hidden bg-cover bg-no-repeat">

    <img
      alt=""
      class="absolute h-full w-screen z-0 object-cover"
      src="@/assets/images/pages/background.jpg"
    >

    <slot />
  </div>
</template>

<script>
  export default {
    name: "LayoutImageBackground"
  }
</script>

<style scoped>

</style>