<template>
  <div>
    <input class="border-2 border-red-500 m-2" type="text" placeholder="ことば" v-model="inputText" />
  </div>
  <div>
    <input class="border-2 border-red-500" type="text" placeholder="意味" v-model="inputMean" />
  </div>
  <button class="m-2 btn-primary" @click="addInputText">追加</button>
</template>
<script>
export default {
  emits: ["input-text"],
  data: () => ({
    inputText: null,
    inputMean: null,
  }),
  methods: {
    addInputText() {
      this.$emit("input-text", this.inputText, this.inputMean);
      this.inputText = "";
      this.inputMean = "";
    },
  },
};
</script>
