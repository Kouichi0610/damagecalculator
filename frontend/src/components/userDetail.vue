// storeサンプルお試し
<template>
    <div class="container">
        <div v-if="profile.user">
            <p>first: {{first}}</p>
            <p>
                FullName: {{ fullName }}
            </p>
            <p>
                Email: {{ email }}
            </p>
        </div>
        <div v-if="profile.error">
            oops an error occured
        </div>
    </div>
    
</template>

<script lang="ts">
import Vue from 'vue';
import {State, Action, Getter} from 'vuex-class';
import Component from 'vue-class-component';
//import {ProfileState, User} from '../store/profile/types';
const namespace: string = 'profile';

@Component
export default class UserDetail extends Vue {
    @State('profile') profile: ProfileState;
    @Action('fetchData', { namespace }) fetchData: any;
    @Getter('fullName', { namespace }) fullName: string;


    mounted() {
        // fetching data as soon as the comonent's been mounted
        this.fetchData();
    }

    get email() {
        const user: User = this.profile && this.profile.user;
        return (user && user.email) || '';
    }
    get first() {
        console.log('error:' + this.profile.error);
        console.log('title:' + this.profile.title);
        const user: User = this.profile && this.profile.user;
        return (user && user.firstName) || '';
    }
}
</script>

<style scoped>
</style>