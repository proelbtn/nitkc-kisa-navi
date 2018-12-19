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
    div.columns
      div.column.is-10-touch.is-offset-1-touch.is-8-desktop.is-offset-2-desktop
        div.cards(v-for="result in results")
          div.box.card
            h5.title.is-5.name {{ result.name }}
            h6.subtitle.is-6.price {{ result.open / 60 }}:{{ zeropad(result.open % 60, 2) }} ~ {{ result.close / 60 }}:{{ zeropad(result.close % 60, 2) }}
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
          active: true,
          title: 'Shops',
          url: '/search/shops'
        },
        {
          active: false,
          title: 'Souvenirs',
          url: '/search/souvenirs'
        }
      ],
      genres: [{ id: 0, name: '---' }],
      results: [],
      name: null,
      genre: 0
    }
  },
  computed: {},
  created() {
    G(this.$axios, '/graphql', {
      query: 'query { shops { genres { id name } } }'
    })
      .then(res => {
        this.genres = res.shops.genres.sort()
        this.genres.unshift({ id: 0, name: '---' })
      })
      .catch(console.error)

    if (navigator.geolocation) {
      navigator.geolocation.getCurrentPosition(
        pos => {},
        error => {
          console.error(error.code)
        }
      )
    }
  },
  methods: {
    zeropad(num, len) {
      return ('0000000000' + num).slice(-len)
    },
    search(event) {
      G(this.$axios, '/graphql', {
        query:
          'query($name: String) { shops { records(name: $name) { id genre_id name address latitude longitude open close } } }',
        variables: { name: this.name }
      })
        .then(resp => {
          this.results = resp.shops.records
        })
        .catch(reason => {
          console.error(reason)
        })
    },
    reset(event) {
      this.name = null
      this.genre = 0
      this.results = []
    }
  }
})
</script>


<style lang="scss">
</style>
