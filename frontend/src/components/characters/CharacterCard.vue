<template>
  <div class="relative w-full h-[200px] flex flex-col group">
    <router-link
      :to="`/characters/${$props.character.id}`"
      class="relative w-full flex-grow overflow-hidden"
    >
      <img
        :src="classLogo"
        alt=""
        class="absolute w-full -top-[75%] object-contain z-10 brightness-[40%] group-hover:brightness-[60%] transition-all"
      >

      <div class="relative z-20 px-4 py-8 h-full flex gap-2">
        <img
          :src="profilePicture"
          alt=""
          class="h-full aspect-square rounded group-hover:scale-110 transition-all"
        >

        <div class="flex justify-center flex-col flex-grow">
          <h3 class="text-gold">{{ $props.character.name }}</h3>
          <p class="text-xs md:text-sm text-gold-light">
            Level {{ $props.character.level }} |
            {{ $props.character.race.name }} |
            {{ $props.character.classes.map(classType => classType.name).join('/') }} /
            {{ $props.character.subclasses.map(subclass => subclass.name).join('/') }}
          </p>
          <p
            v-if="$props.character.user"
            class="text-xs md:text-sm text-gold-light"
          >Player: {{ $props.character.user.username }}</p>
        </div>

      </div>
    </router-link>

    <div class="relative character-card-border w-full px-8 py-2 flex-between gap-2 border border-gray-400 bg-white">
      <h6>
        <button class="px-4 py-2 uppercase rounded text-ink font-bold transition-all bg-white hover:bg-gray-200">View</button>
      </h6>
      <h6>
        <button class="px-4 py-2 uppercase rounded text-ember font-bold transition-all bg-white hover:bg-ember-light hover:text-gold-light">Delete</button>
      </h6>
    </div>
  </div>
</template>

<script>
  import FileService from "@/services/file.service";
  import defaultProfilePicture from "@/assets/images/pages/characters/default-avatar.png"

  export default {
    name: "CharacterCard",
    props: {
      character: {
        type: Object,
        required: true
      }
    },
    data() {
      return {
        profilePicture: "",
        classLogo: "",
      }
    },
    mounted() {
      this.getProfilePicture()
      this.getClassLogo()
    },
    methods: {
      async getProfilePicture() {
        const location = this.$props.character.profile_picture?.location
        if (!location) {
          this.profilePicture = defaultProfilePicture
          return
        }

        const res = await FileService.get(location)
        this.profilePicture = URL.createObjectURL(res.data)
      },
      async getClassLogo() {
        const location = this.$props.character.classes[0].logo?.location
        if (!location) return

        const res = await FileService.get(location)
        this.classLogo = URL.createObjectURL(res.data)
      }
    }
  }
</script>

<style scoped>
.character-card-border::after {
  @apply content-[''] z-20 block absolute bottom-[-4px] left-0 h-[12px] w-full bg-no-repeat bg-cover;
  background-image: url("@/assets/images/pages/characters/character-card-border.png")
}
</style>