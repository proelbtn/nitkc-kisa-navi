<template lang="pug">
  section
    div
      div.field
        label.label Name
        div.control
          input.input(type='text', placeholder='input the name' v-model='name')
      div.field
        div.control
          button.button.is-link(v-on:click='search()') Search
</template>

<script lang="ts">
import Vue from 'vue'
import G from '../middleware/graphql-request-wrapper'

export default Vue.extend({
  data() {
    return {
      name: ''
    }
  },
  methods: {
    search: function(event) {
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
