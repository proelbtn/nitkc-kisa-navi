<template lang="pug">
  main.container
    div.columns
      div.column.is-10-touch.is-offset-1-touch.is-8-desktop.is-offset-2-desktop
        div.content
          h1.title Search
        SearchBoard(:specs="specs", v-on:click-search-button="search()" v-on:click-reset-button="reset()")
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

        hr

        div.box
        div.icon
          <i class="fal fa-abacus"></i>
</template>

<script lang="ts">
import Vue from 'vue'
import SearchBoard from '../../components/SearchBoard.vue'
import G from '../../middleware/graphql-request-wrapper'

export default Vue.extend({
  components: { SearchBoard },
  data() {
    return {
      specs: [
        {
          active: false,
          title: 'Foods',
          url: '/search/foods'
        },
        {
          active: true,
          title: 'Shops',
          url: '/search/shops'
        },
        {
          active: false,
          title: 'Souvenirs',
          url: '/search/souvenirs'
        }
      ],
      genres: [{ id: 0, name: '---' }],
      name: null,
      genre: 0,
      latitude: 35.3816,
      longitude: 139.9262
    }
  },
  computed: {},
  created() {
    G(this.$axios, '/graphql', {
      query: 'query { shops { genres { id name } } }'
    })
      .then(res => {
        this.genres = res.shops.genres.sort()
        this.genres.unshift({ id: 0, name: '---' })
      })
      .catch(console.error)

    if (navigator.geolocation) {
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
    reset(event) {
      this.name = null
      this.genre = 0
    }
  }
})
</script>


<style lang="scss">
</style>
