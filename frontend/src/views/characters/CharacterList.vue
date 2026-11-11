<template>
  <div class="py-8 md:py-16">
    <div
      v-if="!loading"
      class="h-fit w-full grid grid-cols-1 sm:grid-cols-2 lg:grid-cols-3 gap-8 md:gap-16"
    >
      <character-card
        v-for="character in characters"
        :key="`character-card-${character.id}`"
        :character="character"
      />
    </div>
  </div>
</template>


<script>
  import debounce from "lodash/debounce";
  import HelperService from "@/services/helper.service";
  import CharacterService from "@/services/character.service";
  import CharacterCard from "@/views/characters/sections/CharacterCard.vue"

  export default {
    name: "CharacterList",
    components: {CharacterCard},
    data() {
      return {
        loading: true,
        characters: [],
        filters: {
          page: 0,
          page_size: 24,
          search: "",
        },
        totalCount: 0,
        debouncedSearch: null,
      }
    },
    mounted() {
      this.debouncedSearch = debounce(this.getMoreCharacters, 500);
      this.getCharacters()
    },
    methods: {
      getMoreCharacters(page) {
        this.filters.page = page;
        this.getCharacters()
      },
      async getCharacters() {
        this.loading = true;

        try {
          const res = await CharacterService.list(this.filters);

          this.characters = res.data.data || [];
          this.filters.page = res.data.meta.page;
          this.filters.page_size = res.data.meta.page_size;
          this.totalCount = res.data.meta.total_count;
        } catch (error) {
          let errorText = "Could not get characters, please refresh and try again";

          if (res && res.data.error) {
            errorText = res.data.error;
          }

          HelperService.errorToast(this.$toast, errorText);
        } finally {
          this.loading = false;
        }
      }
    }
  }
</script>

<style scoped>

</style>