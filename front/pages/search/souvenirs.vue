<template lang="pug">
  main.container
    div.columns
      div.column.is-10-touch.is-offset-1-touch.is-8-desktop.is-offset-2-desktop
        div.content
          h1.title Search
        SearchTab(:specs="specs")
        div.content
          div.field
            label.label Name
            div.control
              input.input(type="text", name="name", v-model="name")
          div.field
            label.label Genre
            div.control
              div.select
                select(v-model="genre")
                  option(v-for="genre in genres" v-bind:value="genre.id") {{ genre.name }}
          div.field
            label.label Latitude
            div.control
              input.input(type="text", placeholder="...", disabled, v-model="latitude")
          div.field
            label.label Longitude
            div.control
              input.input(type="text", placeholder="...", disabled, v-model="longitude")
          div.field
            button.button.is-link(v-on:click="search()") Search

        hr

        div.box
        div.icon
          <i class="fal fa-abacus"></i>
</template>

<script lang="ts">
import Vue from 'vue'
import SearchTab from '../../components/SearchTab.vue'
import G from '../../middleware/graphql-request-wrapper'

export default Vue.extend({
  components: { SearchTab },
  data() {
    return {
      specs: [
        {
          active: false,
          title: 'Foods',
          url: '/search/foods'
        },
        {
          active: false,
          title: 'Shops',
          url: '/search/shops'
        },
        {
          active: true,
          title: 'Souvenirs',
          url: '/search/souvenirs'
        }
      ],
      genres: [],
      name: null,
      genre: 1,
      latitude: 35.3816,
      longitude: 139.9262
    }
  },
  computed: {},
  created() {
    G(this.$axios, '/graphql', {
      query: 'query { souvenirs { genres { id name } } }'
    })
      .then(res => {
        this.genres.foods = res.foods.genres.sort()
      })
      .catch(console.error)

    if (navigator.geolocation) {
      // TODO: Loading Animation
      navigator.geolocation.getCurrentPosition(
        pos => {
          this.latitude = pos.coords.latitude
          this.longitude = pos.coords.longitude
        },
        error => {
          console.error(error.code)
        }
      )
    }
  },
  methods: {
    search(event) {
      console.log(this.name, this.genre, this.latitude, this.longitude)
      /*
      G(this.$axios, '/graphql', {
        query: 'query($name: String) { foods(name: $name) { name } }',
        variables: { name: this.name }
      })
        .then(console.log)
        .catch(reason => {
          console.error(reason)
        })
        */
    },
    changeFocuses(mode) {
      this.name = ''
      this.genre = 0
      this.focuses = mode
    }
  }
})
</script>


<style lang="scss">
</style>
