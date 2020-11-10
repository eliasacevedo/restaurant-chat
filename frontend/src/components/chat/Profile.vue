<template>
    <section class="profile">
        <figure class="profile--photo">
            <img :src="user.photoPath" class="profile--photo--img" />
        </figure>
        <h2 class="profile--name">{{user.name}}</h2>
    </section>
    
</template>

<script lang="ts">
import { User } from '@/models/user'
import { Component, Vue, Prop } from 'vue-property-decorator'

@Component({ name: 'Profile' })
export default class Profile extends Vue{
  @Prop() user!: User;
}
</script>

<style lang="scss" scoped>
    .profile{
        display: flex;
        align-items: center;
        height: 100%;
        width: 100%;
        padding-left: 10px;
        padding-right: 10px;

        &--photo{
            width: 50px;
            height: 50px;
            border-radius: 100%;
            overflow: hidden;
            display: flex;
            justify-content: center;
            align-items: center;
            margin-right: 5px;
            border: 1px solid #efefef;
            &--img{
                width: 100%;   
            }
        }
    }
</style>
