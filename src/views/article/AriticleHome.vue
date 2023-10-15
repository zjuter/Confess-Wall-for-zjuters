<script setup>
import { artGetDetailService } from "@/api/article";
import { ref, onMounted } from 'vue';

const articles = ref([]); // 存储多篇文章内容

onMounted(async () => {
  let id = 10; // 起始ID
  const ids = [];
  while (true) {
    const response = await artGetDetailService(id); // 获取文章详细信息
    console.log(response);
    if (response.data.data.id) {
      ids.push(id); // 如果文章存在，则将ID加入数组
    } else {
      break; // 如果文章不存在，终止循环
    }
    id++;
  }

  const promises = ids.map(id => artGetDetailService(id)); // 创建 Promise 对象
  const responses = await Promise.all(promises); // 等待所有 Promise 完成
  console.log(responses);
  articles.value = responses.map(res => ({
    pub_date: res.data.data.pub_time,
    cover_img: res.data.data.cover_img,
    id: res.data.data.id,
    username: res.data.data.cate_name === "匿名" ? "匿名" : res.data.data.title,
    content: res.data.data.content
  }));
});
</script>


<template>
  <page-container title="表白墙">
    <!-- 使用 v-for 指令循环渲染多篇文章 -->
    <div v-for="article in articles" :key="article.id" class="article-block">
      <el-col
          v-for="(o, index) in 1"
          :key="o"
          :span="8"
          :offset="index > 0 ? 1 : 0"
      >
        <el-card :body-style="{ padding: '0px' }">
<!--          <img :src="'http://127.0.0.1:3306' + article.cover_img" class="image" />-->
          <div style="padding: 14px">
            <div class="top">
              <span>{{ article.username }}</span>
              <p class="time">{{ article.pub_date }}</p>
            </div>
            <div class="bottom">
              <p class="content" v-html="article.content"></p>
            </div>
          </div>
        </el-card>
      </el-col>
    </div>
  </page-container>
</template>

<style lang="scss" scoped>
.article-block {
  margin-bottom: 20px; // 添加块之间的下边距
}

.article-username {
  font-weight: bold; // 设置用户名字体加粗
}

.article-content {
  margin-top: 5px; // 设置内容与用户名之间的上边距
}

.time {
  font-size: 12px;
  color: #999;
}

.content {
  margin-top: 5px;
  white-space: pre-wrap; /* 设置文本换行 */
}

.top {
  display: flex;
  justify-content: space-between;
  align-items: center;
  margin-bottom: 10px;
}

.bottom {
  line-height: 12px;
}

.button {
  padding: 0;
  min-height: auto;
}

.image {
  width: 100%;
  display: block;
}
</style>
