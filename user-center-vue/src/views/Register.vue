<template>
  <div id="register-page">
    <div id="register-card">
      <p class="title">注册</p>
      <a-form :model="formState" @finish="onFinish" @finishFailed="onFinishFailed"
              :label-col="{ span: 5 }"
              :wrapper-col="{ span: 19 }">

        <a-form-item
            label="账号"
            name="username"
            :rules="[{ required: true, message: 'Please input your username!' }]"
        >
          <a-input v-model:value="formState.username"/>
        </a-form-item>

        <a-form-item
            label="密码"
            name="password"
            :rules="[{ required: true, message: 'Please input your password!' }]"
        >
          <a-input-password v-model:value="formState.password"/>
        </a-form-item>

        <a-form-item :wrapper-col="{ span: 24 }" class="btn-row">
          <a-button type="primary" html-type="submit">注册</a-button>
          <br>
        </a-form-item>
      </a-form>
    </div>
  </div>
</template>


<script setup>
import {reactive} from 'vue';
import {message} from 'ant-design-vue';
import {useRouter} from "vue-router";
import {register} from '../api/user';

const formState = reactive({
  username: '',
  password: '',
  remember: true,
});
const router = useRouter();

const onFinish = async(values) => {
  try {
    const res = await register({
      username: values.username,
      password: values.password,
    });
    const msg = res.data.msg;

    if (msg.includes('成功')) {
      message.success(msg);
      router.push('/login');//注册成功后跳到登录页面
    } else {
      message.warning(msg);
    }
  } catch (e) {
    console.error(e);
    message.error('注册请求失败，检查后端');
  }
};

const onFinishFailed = (errorInfo) => {
  console.log('Failed:', errorInfo);
};
</script>


<style scoped>
#register-page {
  display: flex;
  justify-content: center;
  box-sizing: border-box;
}

#register-card {
  width: 100%;
  max-width: 400px;
  background: #fff;
  border-radius: 20px;
  box-shadow: 0 10px 30px rgba(0, 0, 0, 0.05);
  padding: 24px 28px 32px;
  margin-top: 30px;
}

#register-card .title {
  text-align: center;
  padding-bottom: 16px;
  font-size: 30px;
}

.btn-row {
  text-align: center;
}
</style>