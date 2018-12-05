<template lang="pug">
  main.columns.container
    div.column.is-10-touch.is-offset-1-touch.is-8-desktop.is-offset-2-desktop
      div.content
        h1.title hogehoge
        p
          | fugafugafugafuga
      div.content
        h2.subtitle リンク
        div.field.is-grouped
          p.control
            nuxt-link.button.is-info.is-rounded(to="/about") about
          p.control
            nuxt-link.button.is-primary.is-rounded(to="/search") search
</template>

<script lang="ts">
</script>

<style lang="scss">
</style>
