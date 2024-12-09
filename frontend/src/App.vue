<template>
  <div id="app">
    <Leaderboard />
  </div>
</template>

<script lang="ts">
import Leaderboard from './components/My_Leaderboard.vue'

export default {
  components: {
    Leaderboard,
  },
}
</script>

<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  text-align: center;
}
</style>
