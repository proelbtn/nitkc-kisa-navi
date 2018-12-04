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
import axios from 'axios'

export default Vue.extend({
  data: function() {
    return {
      name: ''
    }
  },
  methods: {
    search: function(event) {
      this.$axios
        .post('/graphql', {
          query: 'query { foods(name: "' + this.name + '") { name } }'
        })
        .then(console.log)
    }
  }
})
</script>


<style lang="scss">
</style>
