<template>
    <router-view/>
</template>
<script>
export default {
    name: 'App'
};
</script>
<style lang="scss">
@import '~/styles/main.scss';
@import url('https://rsms.me/inter/inter.css');

html {
    font-family: 'Inter', sans-serif;
}
@supports (font-variation-settings: normal) {
  html {
      font-family: 'Inter var', sans-serif;
  }
}

body {
    margin: 0;
    -webkit-overflow-scrolling: touch;
}

body,
html,
#app,
.app__container {
    height: 100%;
}
</style>
