<template>
  <div id="main">
    <div class="row">
      <div class="input-field col s12">
        <input v-model="videoID" id="email" type="email" class="validate" />
        <label for="email">Video ID</label>
      </div>
    </div>
    <button @click="sendToBackend()" class="waves-effect waves-light btn">
      <i class="material-icons left">search</i>Search
    </button>
    <template v-if="title !== null && ratio !== null">
      <div class="result">
        <p>
          <span>『{{title}}』</span>のコメントは
          <span>{{ratio}}%</span>が海外のコメントです
        </p>
      </div>
    </template>
    <div class="ranking-wrapper collection">
      <a class="collection-item" v-for="(video, id) in videos" :key="id">
        <a
          class="teal-text text-lighten-2"
          :href="'https://www.youtube.com/watch?v=' + video.video_id"
        >
          <span class="badge">{{video.ration}}%</span>
          {{video.title}}
        </a>
      </a>
    </div>
  </div>
</template>
  
<script>
import axios from "axios";
export default {
  data() {
    return {
      videoID: "",
      title: null,
      ratio: null,
      videos: []
    };
  },
  mounted() {
    axios.get("/videos").then(res => {
      for (let i = 0; i < res.data.length; i++) {
        this.videos.push(res.data[i]);
      }
      console.log(this.videos);
    });
  },
  // updated() {
  //   this.videos = [];
  //   axios.get("/videos").then(res => {
  //     for (let i = 0; i < res.data.length; i++) {
  //       this.videos.push(res.data[i]);
  //     }
  //     fmt.Println(this.videos);
  //   });
  // },
  methods: {
    sendToBackend() {
      const params = new URLSearchParams();
      params.append("query", this.videoID);
      axios.post("/video", params).then(res => {
        this.title = res.data.title;
        this.ratio = res.data.ration;
      });
    }
  }
};
</script>

<style lang="scss">
#main {
  width: 50%;
  margin: auto;
  margin-top: 12%;
  .result {
    margin-top: 10%;
    span {
      font-weight: bold;
      font-size: 1.5rem;
    }
  }
  .ranking-wrapper {
    width: 300px;
    position: absolute;
    top: 12%;
    left: 30px;
  }
}
</style>