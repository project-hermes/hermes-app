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

body {
    margin: 0;
    -webkit-overflow-scrolling: touch;
}

* {
    font-family: Arial, Helvetica, sans-serif;
}

body,
html,
#app,
.app__container {
    height: 100%;
}
</style>
