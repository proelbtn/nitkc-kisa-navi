<template lang="pug">
  div.tabs.is-centered
    ul
      li(v-for="spec in specs" v-bind:class="{ 'is-active': spec.active }")
        nuxt-link(v-bind:to="spec.url") {{ spec.title }}
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
