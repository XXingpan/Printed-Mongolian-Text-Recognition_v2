<template>
    <div class="login_wrap">
        <!-- 在顶部添加系统标识 -->
        <div class="system-logo">
            <!-- <img src="/path/to/logo.png" alt="印刷蒙古文文本识别系统" /> -->
            <h1>印刷蒙古文文本识别系统</h1>
        </div>

        <a-form ref="formRef" :model="loginData" @submit.prevent="submitForm" class="form_wrap" layout="vertical">
            <a-form-item label="用户名" name="username" :rules="usernameRules">
                <a-input v-model="loginData.username" placeholder="请输入用户名" />
            </a-form-item>

            <a-form-item label="密码" name="password" :rules="passwordRules">
                <a-input-password v-model="loginData.password" placeholder="请输入密码" />
            </a-form-item>

            <a-form-item>
                <a-button type="primary" html-type="submit" :loading="loading" block>
                    登录
                </a-button>
                <!-- <a class="forgot-password" href="/forgot-password">忘记密码?</a> -->
            </a-form-item>

            <div class="extra-links">
                没有账户？<a href="/register">注册新账户</a>
            </div>
        </a-form>
    </div>
</template>

<script>
import Vue from 'vue';
import axios from 'axios';

export default Vue.extend({
    name: 'Login',
    data() {
        return {
            loginData: {
                username: '',
                password: ''
            },
            loading: false,
            errors: {
                username: null,
                password: null
            },
            usernameRules: [
                { required: true, message: '请输入用户名', trigger: 'blur' }
            ],
            passwordRules: [
                { required: true, message: '请输入密码', trigger: 'blur' }
            ]
        };
    },
    methods: {
        validateFields() {
            this.errors.username = this.loginData.username ? null : '请输入用户名';
            this.errors.password = this.loginData.password ? null : '请输入密码';
            return !this.errors.username && !this.errors.password;
        },
        submitForm() {
            if (this.validateFields()) {
                this.loading = true;
                const requestData = JSON.parse(JSON.stringify(this.loginData)); // 清理数据
                // console.log('Sending data2:', JSON.stringify(requestData));

                axios.post('/api/sign-in', requestData, {
                    headers: {
                        'Content-Type': 'application/x-www-form-urlencoded'
                    }
                }).then(response => {
                    // console.log('Response:', response);
                    if (response.data.message) {
                        //alert(response.data.message);
                        // 从后端响应中获取userID
                        const userID = response.data.userID;
                        // 使用路由导航传递userID参数
                        // 在跳转前存储 userID
                        sessionStorage.setItem('userID', userID);
                        this.$router.push({
                            path: '/index',
                            // query: { userID: userID }
                        });
                    }
                }).catch(error => {
                    // console.error('Error response:', error.response);
                    if (error.response && error.response.data.error) {
                        alert(error.response.data.error);
                    } else {
                        alert('登录失败，请稍后再试');
                    }
                }).finally(() => {
                    this.loading = false;
                });
            } else {
                console.log('验证失败');
            }
        }
    }
});
</script>

<style scoped>
.login_wrap {
    display: flex;
    flex-direction: column;
    /* 更新为垂直布局 */
    justify-content: center;
    align-items: center;
    height: 100vh;
    background: #2d3761;
}

.system-logo {
    text-align: center;
    margin-bottom: 20px;
}

.system-logo img {
    height: 80px;
    /* 根据需要调整图像大小 */
}

.system-logo h1 {
    color: #ffffff;
    /* 使标题颜色与背景相匹配 */
    margin: 0;
    font-size: 20px;
    /* 根据界面风格调整字号 */
}

.form_wrap {
    width: 300px;
    padding: 20px;
    background: #fff;
    border-radius: 8px;
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.1);
}

.forgot-password {
    display: block;
    margin-top: 10px;
    text-align: right;
}

.extra-links {
    margin-top: 15px;
    text-align: center;
}
</style>