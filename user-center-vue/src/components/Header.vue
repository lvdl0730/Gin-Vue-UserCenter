<template>
  <div class="layout">
    <a-row :wrap="false">
      <a-col flex="250px">
        <div class="title-bar">
          <img src="../assets/logo.png" class="logo" alt="logo"/>
          <div class="title-text"><b>小夏的博客</b></div>
        </div>
      </a-col>

      <a-col flex="auto">
        <a-menu v-model:selectedKeys="selectedKeys" mode="horizontal" :items="items" @select="onMenuSelect"/>
      </a-col>

      <a-col flex="80px">
        <div class="button">
          <a-button type="primary" href="/login">登录</a-button>
        </div>
      </a-col>
    </a-row>
  </div>
</template>

<script lang="ts" setup>
import {computed, h, ref} from 'vue';
import {MailOutlined, AppstoreOutlined, SettingOutlined} from '@ant-design/icons-vue';
import {MenuProps} from 'ant-design-vue';

import {useRouter, useRoute} from 'vue-router';

const router = useRouter();
const route = useRoute();

const current = ref<string[]>(['/']);
const items = ref<MenuProps['items']>([
  {
    key: '/',
    icon: () => h(MailOutlined),
    label: '首页',
    title: '首页',
  },
  {
    key: '/login',
    icon: () => h(AppstoreOutlined),
    label: '用户登录',
    title: '用户登录',
  },
  {
    key: '/register',
    icon: () => h(SettingOutlined),
    label: '用户注册',
    title: '用户注册',
  },
  {
    key: '/manage',
    icon: () => h(AppstoreOutlined),
    label: '用户管理',
    title: '用户管理',
  },
  {
    key: '/nav',
    label: h('a', {href: 'https://www.xxc0609.top', target: '_blank'}, '网页导航'),
    title: '网页导航',
  },
]);
const selectedKeys = computed(() => [route.path]);
const onMenuSelect: MenuProps['onSelect'] = ({key}) => {
  router.push(key as string);
  current.value = [key as string];
}
</script>

<style scoped>
.title-bar {
  display: flex;
  align-items: center;
  background: #fff;
  font-size: 23px;
  color: orange;
}
.title-text {
  margin-left: 16px;
}

.logo {
  height: 64px;
}

</style>