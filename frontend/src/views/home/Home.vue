<template>
  <div class="py-8 md:py-16">
    <div class="h-fit w-full grid grid-cols-1 lg:grid-cols-3 gap-8 md:gap-16">
      <home-card
        v-for="card in cards"
        :key="`home-card-${card.title.toLowerCase()}`"
        :body-text="card.bodyText"
        :image-path="card.imagePath"
        :route="card.route"
        :title="card.title"
      />
    </div>
  </div>
</template>

<script>
  import HomeCard from "@/views/home/sections/HomeCard.vue";
  import CharactersImage from "@/assets/images/pages/home/characters.png"

  export default {
    name: "Home",
    components: {HomeCard},
    data() {
      return {
        cards: [
          {
            title: "Characters",
            bodyText: "Create your character in seconds with our digital builder",
            route: "/characters",
            imagePath: CharactersImage,
          },
        ]
      }
    },
  }
</script>

<style scoped />