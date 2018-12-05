<template lang="pug">
  main.columns.container
    div.column.is-10-touch.is-offset-1-touch.is-8-desktop.is-offset-2-desktop
      div.content
        h1.title Search
      div.tabs.is-centered
        ul
          li(v-bind:class="{ 'is-active': this.focuses == 'foods' }")
            a(v-on:click="changeFocuses('foods')") Foods
          li(v-bind:class="{ 'is-active': this.focuses == 'shops' }")
            a(v-on:click="changeFocuses('shops')") Shops
          li(v-bind:class="{ 'is-active': this.focuses == 'souvenirs' }")
            a(v-on:click="changeFocuses('souvenirs')") Souvenirs
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
                option(v-for="genre in genres[focuses]" v-bind:value="genre.id") {{ genre.name }}
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
</template>

<script lang="ts">
import Vue from 'vue'
import G from '../middleware/graphql-request-wrapper'

export default Vue.extend({
  data() {
    return {
      focuses: 'foods',
      genres: {
        foods: [],
        shops: [],
        souvenirs: []
      },
      name: null,
      genre: 0,
      latitude: 35.3816,
      longitude: 139.9262
    }
  },
  created() {
    G(this.$axios, '/graphql', {
      query: 'query { foods { genres { id name } } }'
    }).then(res => {
      this.genres.foods = res.foods.genres.sort()
    })

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
