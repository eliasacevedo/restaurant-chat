<template>
  <v-app>
    <ServicesProvider>
      <Chat />
    </ServicesProvider>
  </v-app>
</template>

<script lang="ts">
  import { Component, Vue } from 'vue-property-decorator';
  import Chat from './pages/Chat.vue'
  import ServicesProvider from '@/providers/ServicesProvider.vue'

  @Component({
    components: {
      Chat,
      ServicesProvider
    },
  })
  export default class App extends Vue {}
</script>

<style lang="scss">
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  background-color: #fefefe;
}
</style>
