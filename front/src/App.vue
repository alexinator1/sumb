<template>
  <div id="app">
    <AdminPanel />
  </div>
</template>

<script>
import AdminPanel from './components/AdminPanel.vue'

export default {
  name: 'App',
  components: {
    AdminPanel
  }
}
</script>

<style>
#app {
  font-family: 'Inter', -apple-system, BlinkMacSystemFont, 'Segoe UI', Roboto, sans-serif;
  min-height: 100vh;
  background-color: #f8fafc;
}
</style>