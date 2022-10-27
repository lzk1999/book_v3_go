<template>
  <el-col :md="24" :lg="11"
    class="dark-transition rounded-md shadow-2xl dark:shadow-none bg-white dark:bg-slate-800 backdrop-filter bg-opacity-50 dark:bg-opacity-50 mr-4 mb-6">
    <el-row class="bookbox" @mouseenter="bottom_bar_hover = true" @mouseleave="bottom_bar_hover = false">
      <el-col class="bookimg" :span="8">
        <img
          :src="book.img_url"
          :alt="book.name"
          @click="$router.push('/book/' + book.book_id)"
          loading="lazy"
          class="cursor-pointer"
        />
      </el-col>
      <el-col class="bookinfo" :span="14">
        <div class="bookname">
          <span @click="$router.push('/book/' + book.book_id)">{{ book.name }}</span>
        </div>
        <div class="author">作者：{{ book.author }}</div>
        <div class="lastread">简介：{{ book.resume.substring(0,43) }}</div>
        <div>
          <el-row>
            <button class="mybtn" @click="$router.push('/book/' + book.book_id)">开始阅读</button>
          </el-row>
        </div>
      </el-col>

    </el-row>
    <div class="h-0.5 bg-gradient-to-r from-yellow-500 via-pink-500 to-cyan-500"
      :class="{ 'w-full': bottom_bar_hover, 'w-0': !bottom_bar_hover }" style="transition: width 0.1s ease-in-out 0s;">
    </div>
  </el-col>
</template>

<script>
export default {
  name: "bookshelfhot",
  data() {
    return {
      bottom_bar_hover: false,
    }
  },
  props: ["book"],
  emits: ["delBook"],
  methods: {
    select_book(book_id) {
      console.log(book_id);
    },
  },
};
</script>

<style scoped>
.bookbox {
  animation: fadeInDown 0.5s;
}

.bookbox>div {
  padding: 0.5rem;
}

.bookimg img {
  border: none;
  border-radius: .6rem;
  width: 100%;
  cursor: pointer;
  aspect-ratio: 3 / 4;
}

.bookinfo {
  text-align: left;
  display: flex;
  flex-direction: column;
  justify-content: space-around;
}

.bookname span {
  font-size: 1.5rem;
  cursor: pointer;
}

.del_cont {
  align-self: flex-start;
}

.del {
  line-height: 0;
  cursor: pointer;
}

.del svg {
  height: 1.3rem;
  float: right;
}

.el-button {
  background-color: rgba(202, 202, 202, 1);
}

.el-button:hover {
  background-color: #626aef;
}
</style>
