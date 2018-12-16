<template lang="pug">
  div
    div.tabs.is-centered
      ul
        li(v-for="spec in specs" v-bind:class="{ 'is-active': spec.active }")
          nuxt-link(v-bind:to="spec.url") {{ spec.title }}
    div.content
      slot
      div.field
        button.button.is-link(v-on:click="$emit('click-search-button')") Search
        button.button.is-error(v-on:click="$emit('click-reset-button')") Reset
</template>

<script lang="ts">
import Vue from 'vue'
export default Vue.extend({
  props: {
    specs: {
      type: Array,
      default() {
        return []
      }
    }
  }
})
</script>
