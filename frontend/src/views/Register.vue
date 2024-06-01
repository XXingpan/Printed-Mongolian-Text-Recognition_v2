<template>
    <div class="register_wrap">
        <div class="system-logo">
            <h1>印刷蒙古文文本识别系统</h1>
        </div>
        <a-form ref="formRef" :model="registerData" @submit.native.prevent="submitForm" class="form_wrap"
            layout="vertical">
            <a-form-item label="用户名" name="username" :rules="usernameRules">
                <a-input v-model="registerData.username" placeholder="请输入用户名" />
            </a-form-item>

            <a-form-item label="邮箱" name="email" :rules="emailRules">
                <a-input v-model="registerData.email" placeholder="请输入邮箱" />
            </a-form-item>

            <a-form-item label="密码" name="password" :rules="passwordRules">
                <a-input-password v-model="registerData.password" placeholder="请输入密码" />
            </a-form-item>

            <a-form-item label="确认密码" name="confirmPassword" :rules="confirmPasswordRules">
                <a-input-password v-model="registerData.confirmPassword" placeholder="请确认密码" />
            </a-form-item>

            <a-form-item>
                <a-button type="primary" html-type="submit" :loading="loading" block>
                    注册
                </a-button>
            </a-form-item>

            <div class="extra-links">
                已有账户？<a href="/">登录</a>
            </div>
        </a-form>
    </div>
</template>

<script>
import Vue from 'vue';
import axios from 'axios';

export default Vue.extend({
    name: 'Register',
    data() {
        return {
            registerData: {
                username: '',
                email: '',
                password: '',
            },
            loading: false,
            errors: {
                username: null,
                email: null,
                password: null,
                confirmPassword: null,
            },
            usernameRules: [
                { required: true, message: '请输入用户名', trigger: 'blur' },
                { min: 3, max: 20, message: '用户名长度应为3到20个字符', trigger: 'blur' },
            ],
            emailRules: [
                { required: true, message: '请输入邮箱', trigger: 'blur' },
                { pattern: /^[\w.-]+@([\w.-]+\.)+[\w-]{2,4}$/, message: '请输入有效的邮箱地址', trigger: 'blur' },
            ],
            passwordRules: [
                { required: true, message: '请输入密码', trigger: 'blur' },
                { min: 6, message: '密码长度至少为6个字符', trigger: 'blur' },
            ],
            confirmPasswordRules: [
                { required: true, message: '请确认密码', trigger: 'blur' },
                { validator: this.checkPasswordsMatch, trigger: 'blur' },
            ],
        };
    },
    methods: {
        checkPasswordsMatch(rule, value, callback) {
            if (value !== this.registerData.password) {
                callback(new Error('两次输入的密码不匹配!'));
            } else {
                callback();
            }
        },
        validateFields() {
            this.errors.username = this.registerData.username ? null : '请输入用户名';
            this.errors.email = /^[\w.-]+@([\w.-]+\.)+[\w-]{2,4}$/.test(this.registerData.email) ? null : '请输入有效的邮箱地址';
            this.errors.password = this.registerData.password ? null : '请输入密码';
            this.errors.confirmPassword = (this.registerData.confirmPassword === this.registerData.password) ? null : '两次输入的密码不匹配!';

            return !this.errors.username && !this.errors.email && !this.errors.password && !this.errors.confirmPassword;
        },
        submitForm() {
            if (this.validateFields()) {
                this.loading = true;
                const requestData = JSON.parse(JSON.stringify(this.registerData));
                // console.log('==============Sending data:', JSON.stringify(requestData));
                axios.post('/api/sign-up', requestData, {
                    headers: {
                        'Content-Type': 'application/x-www-form-urlencoded',
                    },
                }).then(response => {
                    if (response.data.userID) {
                        const userID = response.data.userID;
                        // 在跳转前存储 userID
                        sessionStorage.setItem('userID', userID);
                        alert(response.data.message);
                        this.$router.push({
                            path: '/index',
                            // query: { userID: userID }
                        });
                    }
                }).catch(error => {
                    console.error('Error response:', error.response);
                    if (error.response) {
                        if (error.response.status === 400) {
                            alert('无效的请求参数');
                        } else if (error.response.status === 409) {
                            alert('用户名或邮箱已存在');
                        } else if (error.response.data.error) {
                            alert(error.response.data.error);
                        } else {
                            alert('注册失败，请稍后再试');
                        }
                    } else {
                        alert('注册失败，请稍后再试');
                    }
                }).finally(() => {
                    this.loading = false;
                });
            } else {
                console.log('注册失败');
            }
        }
    }
});
</script>

<style scoped>
.register_wrap {
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

.extra-links {
    margin-top: 15px;
    text-align: center;
}
</style>
