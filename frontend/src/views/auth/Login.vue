<template>
  <div class="relative flex-center mx-auto w-full xs:w-[75%] lg:w-[50%] xl:w-[40%]">
    <img
      alt=""
      class="absolute w-full h-full z-10"
      src="@/assets/images/svgs/3x2.svg"
    >

    <div class="relative w-full h-full flex flex-col gap-4 md:gap-8 z-20 px-8 md:px-16 py-8">
      <h1 class="text-gold text-center">LOGIN</h1>
      <validation-observer
        ref="observer"
        v-slot="{handleSubmit}"
        class="flex-grow flex-between flex-col"
      >

        <form
          autocomplete="off"
          class="flex-between flex-grow flex-col"
          @submit.prevent="handleSubmit(login)"
        >
          <section class="w-full flex flex-col gap-6">
            <validation-provider
              v-slot="validationContext"
              :rules="{required: true}"
              class="validation-provider"
              name="Username"
            >
              <input-wrapper
                :value="user.username"
                autocomplete="current-password"
                label="Username"
                name="login-username"
                placeholder="Username"
                @input="user.username = $event"
              />

              <small class="text-ember">{{ validationContext.errors[0] }}</small>
            </validation-provider>

            <validation-provider
              v-slot="validationContext"
              :rules="{required: true}"
              class="validation-provider"
              name="Password"
            >
              <input-wrapper
                :type="passwordFieldType"
                :value="user.password"
                label="Password"
                name="login-password"
                placeholder="Password"
                @input="user.password = $event"
              >
                <feather-icon
                  :icon="passwordToggleIcon"
                  class="cursor-pointer text-ink"
                  size="24"
                  @click="togglePasswordVisibility"
                />
              </input-wrapper>

              <small class="text-ember">{{ validationContext.errors[0] }}</small>
            </validation-provider>
          </section>

          <button
            class="w-3/4 md:w-1/2 px-8 py-2 mt-8 rounded-xl  md:text-2xl mx-auto bg-ember transition-all shadow-lg shadow-transparent text-gold-light hover:bg-ember-light hover:shadow-ink"
            type="submit"
          >
            Login
          </button>
        </form>
      </validation-observer>

      <h4 class="flex-center flex-col md:flex-row text-center text-gold-light w-full">
        Don't have an account?&nbsp;
        <router-link
          class="text-ember transition-all hover:text-ember-light"
          to="/register"
        >
          Register Now
        </router-link>
      </h4>
    </div>
  </div>
</template>

<script>
  import AuthService from "@/services/auth.service";
  import InputWrapper from "@/components/forms/InputWrapper.vue";
  import {togglePasswordVisibility} from "@core/mixins/ui/forms";
  import FeatherIcon from "@core/components/feather-icon/FeatherIcon.vue";
  import HelperService from "@/services/helper.service";

  export default {
    name: "Login",
    components: {InputWrapper, FeatherIcon},
    mixins: [togglePasswordVisibility],
    data() {
      return {
        user: {
          username: '',
          password: ''
        },
        loading: false
      }
    },
    computed: {
      passwordToggleIcon() {
        return this.passwordFieldType === 'password' ? 'EyeIcon' : 'EyeOffIcon'
      },
    },
    methods: {
      async login() {
        this.loading = true
        try {
          const dto = {
            username: this.user.username,
            password: this.user.password,
          }

          await AuthService.login(dto)
          await this.$router.push("/")
        } catch (err) {
          HelperService.errorToast(this.$toast, err, "There was an error logging in. Please try again.")
        } finally {
          this.loading = false
        }
      },
    }
  }
</script>


<style scoped>

</style>