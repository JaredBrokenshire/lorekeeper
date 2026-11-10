<template>
  <div class="w-full bg-[#232b2f] flex-between">
    <div class="max-w-[1340px] w-full py-2 flex justify-end xs:justify-between items-center flex-grow mx-4 md:mx-8 xl:mx-auto">
      <section class="hidden xs:flex gap-2">
        <router-link to="/">
          <img
            alt="Website Logo"
            class="w-16 aspect-square"
            src="@/assets/images/website-logo.png"
          >
        </router-link>
      </section>

      <section class="flex-between gap-2 md:gap-4">
        <div class="bg-[#636b6f] flex-grow flex-center text-sm md:text-lg text-gold-light text-center px-2 py-1 sm:px-4 sm:py-2 rounded transition-all">Hi, {{ $store.getters["auth/currentUser"]?.username ?? "" }}</div>
        <button
          class="bg-ember rounded text-sm md:text-lg px-2 py-1 sm:px-4 sm:py-2 transition-all shadow-lg shadow-transparent hover:bg-ember-light hover:shadow-ink text-gold-light"
          @click="logout"
        >
          LOGOUT
        </button>
      </section>
    </div>
  </div>
</template>

<script>
  import AuthService from "@/services/auth.service";
  import HelperService from "@/services/helper.service";

  export default {
    name: "PageHeader",
    data() {
      return {
        loading: false,
      }
    },
    methods: {
      async logout() {
        this.loading = true;
        try {
          await AuthService.logout();

          await this.$router.push("/login");
        } catch (error) {
          HelperService.errorToast(this.$toast, "Could not logout, please try again");
        } finally {
          this.loading = false;
        }
      }
    }
  }
</script>

<style scoped>

</style>