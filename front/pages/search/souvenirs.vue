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
            label.label Price
            div.control
              input.input(type="number", name="price", v-model="price")
    div.columns
      div.column.is-10-touch.is-offset-1-touch.is-8-desktop.is-offset-2-desktop
        div.cards(v-for="result in results")
          div.box.card
            h5.title.is-5.name {{ result.name }}
            h6.subtitle.is-6.price {{ result.price }}円
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
          active: false,
          title: 'Shops',
          url: '/search/shops'
        },
        {
          active: true,
          title: 'Souvenirs',
          url: '/search/souvenirs'
        }
      ],
      genres: [{ id: 0, name: '---' }],
      results: [],
      name: null,
      genre: 0,
      price: null
    }
  },
  computed: {},
  created() {
    G(this.$axios, '/graphql', {
      query: 'query { souvenirs { genres { id name } } }'
    })
      .then(res => {
        this.genres = res.souvenirs.genres.sort()
        this.genres.unshift({ id: 0, name: '---' })
      })
      .catch(console.error)
  },
  methods: {
    search(event) {
      G(this.$axios, '/graphql', {
        query:
          'query($name: String) { souvenirs { records(name: $name) { id genre_id name price } } }',
        variables: { name: this.name }
      })
        .then(resp => (this.results = resp.souvenirs.records))
        .catch(reason => {
          console.error(reason)
        })
    },
    reset(event) {
      this.name = null
      this.genre = 0
      this.price = null
      this.results = []
    }
  }
})
</script>


<style lang="scss">
</style>
