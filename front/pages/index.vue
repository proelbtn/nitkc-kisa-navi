<template lang="pug">
  main.container
    div.columns
      div.column.is-10-touch.is-offset-1-touch.is-8-desktop.is-offset-2-desktop
        div.content
          h1.title kisa-navi
          p 木更津市の観光アプリです！
          p 食べ物やお土産、お店の検索をすることができます。
        div.content
          h2.subtitle Links
          div.field.is-grouped(style="margin-bottom: 20px")
            p.control.has-text-centered(style="width: 100%")
              nuxt-link.button.is-primary.is-rounded.is-large(to="/search/foods") Search
</template>

<script lang="ts">
</script>

<style lang="scss">
</style>
