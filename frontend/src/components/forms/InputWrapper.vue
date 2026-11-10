<template>
  <section>
    <label
      v-if="$props.label"
      :for="$props.name"
      class="text-gold-light"
    >
      {{ $props.label }}
    </label>

    <div class="input-wrapper">
      <input
        :id="name"
        :autocomplete="$attrs.autocomplete || 'off'"
        :name="name"
        :placeholder="placeholder"
        :type="type"
        :value="value"
        v-bind="{
          ...$attrs,
          autocomplete: undefined
        }"
        @input="$emit('input', $event.target.value)"
      >

      <div class="relative z-20">
        <slot />
      </div>

      <svg
        class="input-decoration"
        preserveAspectRatio="none"
        viewBox="0 0 344 95"
        xmlns="http://www.w3.org/2000/svg"
      >
        <polygon
          fill="#ABABABDB"
          points="8 93 337 93 337 2 8 2 8 93"
        />
        <path
          d="M344,6.39V4.47h-6.14V0h-2.68s-1.06,1.54-3.91,1.54H12.73C9.88,1.54,8.82,0,8.82,0H6.14V4.47H0V6.39c2.53,0,2.67,4.14,2.67,4.14V81.91S2.53,86,0,86v2H6.14v7H8.82V3.31H335.18V91.69H8.82V95s1.06-1.54,3.91-1.54H331.27c2.84,0,3.9,1.52,3.91,1.54h2.68V88H344V86c-2.53,0-2.67-4.13-2.67-4.13V10.53S341.47,6.39,344,6.39ZM6.27,81.91H4.14V12.12H6.27Zm333.79.48h-2.12V12.61h2.12Z"
          fill="#972e2e"
        />
      </svg>
    </div>
  </section>
</template>

<script>
  export default {
    name: "InputWrapper",
    inheritAttrs: false,
    props: {
      type: {
        type: String,
        default: "text"
      },
      value: {
        type: [String, Number],
        default: '',
      },
      placeholder: {
        type: String,
        default: '',
      },
      name: {
        type: String,
        default: '',
      },
      label: {
        type: String,
        default: '',
      },
      validationRules: {
        type: Object,
        default: null,
      }
    }
  }
</script>

<style scoped>
.input-wrapper {
  @apply relative w-full flex justify-between items-center pe-6 ps-[1vw] xl:text-lg
}

.input-wrapper input {
  @apply relative z-10 w-full h-[3rem] px-[1rem] bg-transparent outline-none text-ink placeholder:text-gold-dark;
}

.input-decoration {
  @apply absolute inset-0 w-full h-full pointer-events-none;
}

.input-decoration polygon, .input-decoration path {
  @apply transition-all;
}

.input-wrapper:focus-within .input-decoration polygon {
  fill: #FFFFFFDB;
}
.input-wrapper:focus-within .input-decoration path {
  @apply fill-ember-light
}
</style>