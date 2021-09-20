<template>
  <ul>
    <li class="m-2" v-for="item in fuga" v-bind:key="item">言葉 : 「{{ item.text }}」 意味 : 「{{item.mean}}」</li>
  </ul>
</template>

<script>
export default {
  name: "TextList",
  props: { fuga : Array },
};
</script>

<style></style>
