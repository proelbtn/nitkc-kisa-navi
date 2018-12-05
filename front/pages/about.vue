<template lang="pug">
  main.columns.container
    div.column.is-10-touch.is-offset-1-touch.is-8-desktop.is-offset-2-desktop
      div.content
        h1.title About
        p
          | fugafugafugafuga
      div.content
        h2.subtitle Links
        div.field.is-grouped
          p.control
            nuxt-link.button.is-info.is-rounded(to="/") Top
          p.control
            nuxt-link.button.is-primary.is-rounded(to="/search") Search
</template>

<script lang="ts">
</script>

<style lang="scss">
</style>
