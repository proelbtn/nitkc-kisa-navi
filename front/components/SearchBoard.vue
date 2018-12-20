<template lang="pug">
  div
    div.tabs.is-centered
      ul
        li(v-for="spec in specs" v-bind:class="{ 'is-active': spec.active }")
          nuxt-link(v-bind:to="spec.url") {{ spec.title }}
    div.content
      slot
      div.buttons-box.field.has-text-centered(style="padding-top: 20px")
        button.button.is-link.is-rounded.is-large(v-on:click="$emit('click-search-button')") Search
        button.button.is-danger.is-rounded.is-large(v-on:click="$emit('click-reset-button')") Reset
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

<style lang="scss">
.buttons-box * {
  margin: 0px 10px;
}
</style>
