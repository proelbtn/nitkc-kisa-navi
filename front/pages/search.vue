<template lang="pug">
  main.columns.container
    div.column.is-10-touch.is-offset-1-touch.is-8-desktop.is-offset-2-desktop
      div.content
        h1.title Search
      div.content
        div.field
          label.label Name
          div.control
            input.input(type="text", name="name")
        div.field
          label.label Genre
          div.control
            div.select
              select
                option test1
                option test2
        div.field
          label.label Latitude
          div.control
            input.input(type="text", placeholder="...", disabled, v-model="latitude")
        div.field
          label.label Longitude
          div.control
            input.input(type="text", placeholder="...", disabled, v-model="longitude")
</template>

<script lang="ts">
import Vue from 'vue'
import G from '../middleware/graphql-request-wrapper'

export default Vue.extend({
  data() {
    return {
      name: '',
      genre: '',
      latitude: 35.3816,
      longitude: 139.9262
    }
  },
  created() {
    // TODO: Fetch Genre

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
      if (navigator.geolocation) {
        navigator.geolocation.getCurrentPosition(console.log, console.error)
      }
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
    }
  }
})
</script>


<style lang="scss">
</style>
