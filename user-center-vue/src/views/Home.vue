<script setup>
import {ContactsOutlined, IdcardOutlined, UserOutlined, GlobalOutlined,} from '@ant-design/icons-vue';
</script>

<template>
  <div id="home-page">
    <div id="home-card">
      <div class="title">
        <b>内容介绍</b>
      </div>
      <table class="table-list">
        <thead>
        <tr>
          <th>图标</th>
          <th>标题</th>
          <th>描述</th>
          <th>按钮</th>
        </tr>
        </thead>
        <tbody>
        <tr>
          <th>
            <UserOutlined/>
          </th>
          <th>用户登录</th>
          <th>已有账号直接登录</th>
          <th>
            <a-button type="link" href="/login">登录</a-button>
          </th>
        </tr>
        <tr>
          <th>
            <ContactsOutlined/>
          </th>
          <th>用户注册</th>
          <th>第一次请先注册</th>
          <th>
            <a-button type="link" href="/register">注册</a-button>
          </th>
        </tr>
        <tr>
          <th>
            <IdcardOutlined/>
          </th>
          <th>用户管理</th>
          <th>管理用户信息</th>
          <th>
            <a-button type="link" href="/manage">用户管理</a-button>
          </th>
        </tr>
        <tr>
          <th>
            <GlobalOutlined/>
          </th>
          <th>导航</th>
          <th>个人博客</th>
          <th><a href="https://www.xxc0609.top">导航</a></th>
        </tr>
        </tbody>
      </table>
    </div>
  </div>

</template>

<style scoped>
#home-page {
  display: flex;
  justify-content: center;
  box-sizing: border-box;
}

#home-card {
  width: 100%;
  max-width: 500px;
  background: #fff;
  border-radius: 20px;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.05);
  padding: 24px 28px 32px;
  margin-top: 30px;
}

#home-card .title {
  text-align: center;
  padding-bottom: 16px;
  font-size: 30px;
}

#home-card .table-list {
  width: 100%;
}

.table-list th, td {
  padding: 16px 20px;
  border-bottom: 1px solid #f0f0f0;
}
</style>