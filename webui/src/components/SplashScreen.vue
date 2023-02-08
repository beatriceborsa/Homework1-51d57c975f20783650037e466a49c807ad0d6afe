<template>
    <div :class="{ loader: true, fadeout: !isLoading }">
        <img :src="image" />
    </div>
</template>

<script>
import image from "../assets/logo.png"
export default {
    name: "SplashScreen",
    props: ["isLoading"],
    data() {
        return {
            image: image
        }
    },
};
</script>

<style>
.loader {
    background-color: #ffffff;
    bottom: 0;
    color: white;
    display: block;
    font-size: 32px;
    left: 0;
    overflow: hidden;
    padding-top: 10vh;
    position: fixed;
    right: 0;
    text-align: center;
    top: 0;
}

.fadeout {
    animation: fadeout 3s forwards;
}

@keyframes fadeout {
    to {
        opacity: 0;
        visibility: hidden;
    }
}
</style>