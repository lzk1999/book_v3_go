<template>
  <div class="app w-full sm:w-3/4 2xl:w-2/3 pt-12 sm:pt-20">
    <Header :headerSetting="headerSetting"></Header>
    <Background></Background>
    <div class="book-list">
      <div class="list-title">
        <span>首页推荐</span>
      </div>
      <el-row>
        <bookshelfhot v-for="book in books" :key="book.book_id" :book="book" @delBook="delbook"></bookshelfhot>
      </el-row>
    </div>
    <div class="footer">
      <div class="h-0.5 bg-gradient-to-r from-yellow-500 via-pink-500 to-cyan-500 w-full"></div>
      <div>
        <span>
          Copyright © 2022 By
          <a
            href="http://www.xcya.cn"
            class="italic font-bold text-xl bg-clip-text linear-gradient inverse text-fill-transparent"
          >Silver alien</a>
        </span>
      </div>
      <div>
        <a href="http://beian.miit.gov.cn/">
          <span>京ICP备2022006333号</span>
        </a>
      </div>
    </div>
  </div>
</template>

<script>
import $ from "jquery";
import BackGround from './common/background.vue'
import { Search } from '@element-plus/icons-vue'
import HomeHeader from './header/searchHeader.vue'
import {useRouter} from "vue-router";
import axios from "axios";
import {ref} from "vue";

export default {
  name: "home",
  components: { BackGround,Search },
  props: {
    headerSetting: {
      type: Object
    },
    homeString: {
      type: String,
      default: 'Home'
    },
    homeHref: {
      type: String,
      default: '/'
    }
  },
  setup() {

    const search_string = ref('')

    return { search_string };
  },
  data() {
    return {
      module_more_show: false,
      books:null,
      headerSetting: {},
    };
  },
  mounted() {
    this.initHeader();
    this.initHot();
  },
  methods:{
    initHot() {
      axios
          .get("/api/indexhot")
          .then((res) => {
            this.books = res.data;
          })
          .catch((err) => {
            ElNotification({ message: '请求失败！请刷新！', type: 'error', duration: 2000 });
            console.log(err);
          });
    },
    initHeader() {
      this.headerSetting = {
        headerSettings:
            [
              {
                type: 'search',
                placeholder: "搜索书籍。",
                clickHandle: this.searchbook
              },
            ],
        userSetting: {
          userLogHandle: this.userLog
        }
      }
    },
    searchbook(s) {
      let search_s = s || this.search_string;
      this.searchloading = true;
      this.search_info = "正在搜索中...";
      this.bookshelf_show = false;
      this.bookbox_list_show = true;
      this.$router.push("/book?s=" + search_s);
      axios
          .post("/api/book/search", {
            search_string: search_s,
          })
          .then((res) => {
            if (res.data.data == ""){
              this.forshow = false;
            }else{
              this.forshow = true;
            }
            this.searchloading = false;
            (document.title = "'" + search_s + "'" + "的搜索结果"),
                (this.search_books = res.data.data);
            if (this.search_books.length) {
              this.search_info =
                  '😁 搜索到与"<b>' +
                  search_s +
                  '</b>"相关的书籍，共' +
                  this.search_books.length +
                  "本。（最多显示100本）";
            } else {
              this.search_info =
                  '😭 暂未找到与"<b>' + search_s + '</b>"相关的书籍。';
            }
          })
          .catch((err) => {
            this.searchloading = true;
            ElNotification({ message: '请求失败！请刷新！', type: 'error', duration: 0 });
            console.log(err);
          });
    },
  }
};
</script>

<style scoped>
@import "../assets/css/main.css";
@import "../assets/css/book.css";
.search:hover {
  background: none !important;
  border-bottom: none !important;
}

.input-with-select {
  opacity: 0.9;
  color: #000;
}

.el-input__icon {
  color: #4642c5;
  font-size: 1.3rem;
}
.input-with-select {
  opacity: 0.8;
  color: #000;
}

.el-input__icon {
  color: #4642c5;
  font-size: 1.3rem;
}

.search-two {
  margin: 1rem auto;
  max-width: 40rem;
}

.search-two .el-input__icon {
  font-size: 1.5rem;
  cursor: pointer;
}

.backpic {
  z-index: -1;
  max-height: 100%;
  overflow: hidden;
}

.backpic img:nth-child(1) {
  top: 0;
  left: 5%;
}

.backpic img:nth-child(2) {
  top: 50%;
  left: -30%;
  opacity: 0.7;
}
.page__x {
  width: 1000px;
  height: 700px;
  /* 居中布局 */
  display: flex;
  justify-content: center;
  align-items: center;
  overflow: hidden;
  /* 设置元素被查看位置的视图 */
  perspective: 1800px;
  /* 背景色（兼容性写法） */
  background: #642b73;
  background: linear-gradient(to bottom, #c6426e, #642b73);
}

/* Popular */
h1 {
  /* 底部外边距 */
  margin-bottom: 30px;
  /* z轴偏移 */
  transform: translateZ(35px);
  /* 字母间距 */
  letter-spacing: -1px;
  /* 字号 */
  font-size: 32px;
  /* 字体粗细 */
  font-weight: 800;
  /* 字体颜色 */
  color: #3e3e42;
}

/* Movies */
h3 {
  /* 底部外边距 */
  margin-bottom: 6px;
  /* z轴偏移 */
  transform: translateZ(25px);
  /* 字号 */
  font-size: 16px;
  /* 字体颜色 */
  color: #eb285d;
}

/* 卡片主容器 */
.cards {
  /* 行内块元素 */
  display: inline-block;
  /* 最小宽度 */
  min-width: 595px;
  /* 内边距 */
  padding: 30px 35px;
  /* 设置元素被查看位置的视图 */
  perspective: 1800px;
  /* 旋转基点 */
  transform-origin: 50% 50%;
  /* 使被转换的子元素保留其 3D 转换 */
  transform-style: preserve-3d;
  /* 圆角 */
  border-radius: 15px;
  /* 文本左对齐 */
  text-align: left;
  /* 背景色 */
  background: #fff;
  /* 投影 */
  box-shadow: 0px 10px 20px 20px rgba(0, 0, 0, 0.17);
}

/* 卡片 */
.card {
  /* 行内块元素 */
  display: inline-block;
  /* 宽 */
  width: 175px;
  /* 高 */
  height: 250px;
  /* 相对定位 */
  position: relative;
  /* 隐藏溢出部分 */
  overflow: hidden;
  /* 设置元素被查看位置的视图 */
  perspective: 1200px;
  /* 使被转换的子元素保留其 3D 转换 */
  transform-style: preserve-3d;
  /* z轴偏移 */
  transform: translatez(35px);
  /* 过渡 */
  transition: transform 200ms ease-out;
  /* 文本居中对齐 */
  text-align: center;
  /* 圆角 */
  border-radius: 15px;
  /* 投影 */
  box-shadow: 5px 5px 20px -5px rgba(0, 0, 0, 0.6);
}

/* 除了最后一个卡片之外的卡片 */
.card:not(:last-child) {
  /* 右侧外边距 */
  margin-right: 30px;
}

/* 卡片的图片 */
.card__img {
  /* 相对定位 */
  position: relative;
  /* 高度 */
  height: 100%;
}

/* 卡片背景 */
.card__bg img {
  bottom: -50px;
  left: -50px;
  position: absolute;
  right: -50px;
  top: -50px;
  /* 旋转基点 */
  transform-origin: 50% 50%;
  transform: translateZ(-50px);
  z-index: 0;
}

/* 幽灵公主 图片 */
.princess-mononoke .card__img {
  top: 14px;
  right: -10px;
  height: 110%;
}

/* 幽灵公主 背景 */
.princess-mononoke .card__bg {
  background: url("@/assets/img/3dr_monobg.jpg") center/cover no-repeat;
}

/* 千与千寻 图片 */
.spirited-away .card__img {
  top: 25px;
}

/* 千与千寻 背景 */
.spirited-away .card__bg {
  background: url("@/assets/img/3dr_spirited.jpg") center/cover no-repeat;
}

/* 哈尔的移动城堡 图片 */
.howl-s-moving-castle .card__img {
  top: 5px;
  left: -4px;
  height: 110%;
}

/* 哈尔的移动城堡 背景 */
.howl-s-moving-castle .card__bg {
  background: url("@/assets/img/3dr_howlbg.jpg") center/cover no-repeat;
}

/* 卡片的文本内容 */
.card__text {
  /* 弹性布局 */
  display: flex;
  /* 主轴为垂直方向 */
  flex-direction: column;
  /* 主轴居中对齐 */
  justify-content: center;
  /* 交叉轴的中点对齐 */
  align-items: center;
  /* 宽 */
  width: 100%;
  /* 高 */
  height: 70px;
  /* 绝对定位 */
  position: absolute;
  /* 堆叠顺序 */
  z-index: 2;
  /* 离底部距离 */
  bottom: 0;
  /* 背景色：渐变 */
  background: linear-gradient(to bottom,
  rgba(0, 0, 0, 0) 0%,
  rgba(0, 0, 0, 0.55) 100%);
}

/* 卡片的标题 */
.card__title {
  /* 底部外边距 */
  margin-bottom: 3px;
  /* 设置左右10px内边距 */
  padding: 0 10px;
  /* 字号 */
  font-size: 18px;
  /* 字体的粗细 */
  font-weight: 700;
  /* 字体颜色 */
  color: #fff;
}
</style>
