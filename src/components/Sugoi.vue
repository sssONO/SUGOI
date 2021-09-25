<template>
  <div id="sugoi">
    <div id="nav" class="tab-area-base">
      <ul class="tab-menu-base">
        <!-- *** -->
        <Title :hoge="title" />
        <InputText @input-text="addText" />
        <TextList :fuga="textList" />
        <!-- *** -->
      </ul>
    </div>
  </div>
</template>

<script>
import InputText from "./InputText.vue";
import Title from "./Title.vue";
import TextList from "./TextList.vue";

export default {
  name: "Sugoi",
  components: {
    // HelloWorld,
    InputText,
    Title,
    TextList,
  },
  data: () => ({
    title: "好きな言葉を入力しよう",
    textList: [],
  }),
  methods: {
    addText(value, value2) {
      const newText = {
        id: new Date().toISOString(),
        text: value,
        mean: value2,
      };
      this.textList.unshift(newText);
      this.saveText();
    },
    saveText: function() {
      localStorage.setItem("textList", JSON.stringify(this.textList));
    },
    loadText: function() {
      this.textList = JSON.parse(localStorage.getItem("textList"));
      if (!this.textList) {
        this.textList = [];
      }
    },
  },
  mounted: function() {
    //サイト始めたらここ
    this.loadText();
  },
};
</script>
<style>
#app {
  font-family: Avenir, Helvetica, Arial, sans-serif;
  -webkit-font-smoothing: antialiased;
  -moz-osx-font-smoothing: grayscale;
  text-align: center;
  color: #2c3e50;
  margin-top: 60px;
}
</style>


