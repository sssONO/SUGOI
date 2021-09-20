<template>
  <div>
    <h1>This is an Yamabiko page</h1>
    <input v-model="message" placeholder="Say Yahho" />
    <button v-on:click="doAddProduct">Send</button>
  </div>
</template>

<script>
import axios from "axios";

export default {
  name: "Yamabiko",
  data: () => ({
    message: "",
  }),
  methods: {
    // async Send() {
    //   const yamabiko = await this.CallYamabikoAPI().then((res) => res.json());
    //   this.yamabiko = yamabiko.message;
    //   window.alert(this.yamabiko);
    // },

    doAddProduct() {
      const params = new URLSearchParams();
      params.append("Name", this.message);
      axios.post("/addProduct", params).then((response) => {
        if (response.status != 200) {
          throw new Error("レスポンスエラー");
        } else {
          // 入力値を初期化する
          this.initInputValue();
        }
      });
    },

    initInputValue() {
      this.productName = "";
    },
  },
};
</script>
