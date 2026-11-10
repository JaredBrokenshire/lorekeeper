<template>
  <div class="relative flex-center mx-auto w-full md:w-[75%] lg:w-[60%] xl:w-[50%]">
    <img
      alt=""
      class="absolute w-full h-full z-10"
      src="@/assets/images/svgs/3x2.svg"
    >

    <div class="relative w-full h-full flex flex-col gap-8 z-20 px-12 py-8">
      <h1 class="text-gold text-center">REGISTER</h1>

      <validation-observer
        ref="observer"
        v-slot="{handleSubmit}"
        class="flex-grow flex-between flex-col"
      >

        <form
          autocomplete="off"
          class="flex-between flex-grow flex-col gap-4"
          @submit.prevent="handleSubmit(register)"
        >
          <div class="w-full flex flex-col sm:flex-row gap-4">
            <section class="validation-provider-container sm:w-1/2">
              <validation-provider
                v-slot="validationContext"
                :rules="{required: true}"
                class="validation-provider"
                name="Username"
              >
                <input-wrapper
                  :value="username"
                  autocomplete="off"
                  label="Username"
                  name="register-username"
                  placeholder="Username"
                  @input="username = $event; usernameError = ''"
                />

                <small class="text-ember">
                  {{ validationContext.errors[0] }}
                  {{ usernameError }}
                </small>
              </validation-provider>
            </section>

            <section class="validation-provider-container sm:w-1/2">
              <validation-provider
                v-slot="validationContext"
                :rules="{required: true}"
                class="validation-provider"
                name="Email"
              >
                <input-wrapper
                  :value="email"
                  autocomplete="off"
                  label="Email"
                  name="register-email"
                  placeholder="Email"
                  type="email"
                  @input="email = $event; emailError = ''"
                />

                <small class="text-ember">
                  {{ validationContext.errors[0] }}
                  {{ emailError }}
                </small>
              </validation-provider>
            </section>
          </div>

          <div class="w-full flex flex-col sm:flex-row gap-4 mb-4">
            <section class="validation-provider-container sm:w-1/2">
              <validation-provider
                v-slot="validationContext"
                :rules="{required: true}"
                class="validation-provider"
                name="Password"
              >
                <input-wrapper
                  :type="passwordFieldType"
                  :value="password"
                  label="Password"
                  name="login-password"
                  placeholder="Password"
                  @input="password = $event; passwordError = ''"
                >
                  <feather-icon
                    :icon="passwordToggleIcon"
                    class="cursor-pointer text-ink"
                    size="24"
                    @click="togglePasswordVisibility"
                  />
                </input-wrapper>

                <small class="text-ember">
                  {{ validationContext.errors[0] }}
                  {{ passwordsMatch ? '' : 'Passwords do not match' }}
                  {{ passwordError }}
                </small>
              </validation-provider>
            </section>

            <section class="validation-provider-container sm:w-1/2">
              <validation-provider
                v-slot="validationContext"
                :rules="{required: true}"
                class="validation-provider"
                name="Confirm Password"
              >
                <input-wrapper
                  :type="passwordFieldType"
                  :value="confirm_password"
                  label="Confirm Password"
                  name="login-confirm-password"
                  placeholder="Confirm Password"
                  @input="confirm_password = $event; passwordError = ''"
                />
                  
                <small class="text-ember">{{ validationContext.errors[0] }}</small>
              </validation-provider>
            </section>
          </div>

          <button
            class="w-3/4 md:w-1/2 px-8 py-2 mt-2 md:mt-4 lg:mt-8 rounded-xl  md:text-2xl mx-auto bg-ember transition-all shadow-lg shadow-transparent text-gold-light hover:bg-ember-light hover:shadow-ink"
            type="submit"
          >
            Register
          </button>
        </form>
      </validation-observer>

      <h4 class="flex-center text-gold-light">
        Already have an account?&nbsp;
        <router-link
          class="text-ember transition-all hover:text-ember-light"
          to="/login"
        >
          Login Now
        </router-link>
      </h4>
    </div>
  </div>
</template>

<script>
  import AuthService from "@/services/auth.service";
  import HelperService from "@/services/helper.service";
  import InputWrapper from "@/components/forms/InputWrapper.vue";
  import {togglePasswordVisibility} from "@core/mixins/ui/forms";
  import FeatherIcon from "@core/components/feather-icon/FeatherIcon.vue";

  export default {
    name: "Register",
    components: {InputWrapper, FeatherIcon},
    mixins: [togglePasswordVisibility],
    data() {
      return {
        username: '',
        email: '',
        password: '',
        confirm_password: '',
        loading: false,
        usernameError: "",
        emailError: "",
        passwordError: "",
      }
    },
    computed: {
      passwordToggleIcon() {
        return this.passwordFieldType === 'password' ? 'EyeIcon' : 'EyeOffIcon'
      },
      passwordsMatch() {
        if (this.password && this.confirm_password) {
          return this.password === this.confirm_password
        }
        return true
      }
    },
    methods: {
      async register() {
        this.loading = true

        try {
          const dto = {
            username: this.username,
            email: this.email,
            password: this.password,
            confirm_password: this.confirm_password,
          }

          await AuthService.register(dto)
          await this.$router.push("/")
        } catch (err) {
          switch (err.response.data.error) {
          case 'Username is already taken':
            this.usernameError = "Username is already taken"
            break;
          case 'Email is already registered':
            this.emailError = "Email is already registered"
            break;
          case 'Passwords do not match':
            this.passwordError = "Passwords do not match"
            break;
          default:
            HelperService.errorToast(this.$toast, err, "There was an error registering your account. Please try again.")
          }
        } finally {
          this.loading = false
        }
      },
    }
  }
</script>

<style scoped>

</style>