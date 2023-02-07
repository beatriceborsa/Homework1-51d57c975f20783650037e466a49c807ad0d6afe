<template>
    <div :class="{ loader: true, fadeout: !isLoading }">
        WELCOME TO WASAPHOTO-BB
    </div>
</template>

<script>
export default {
    name: "SplashScreen",
    props: ["isLoading"]
};
</script>

<style>
.loader {
    background-color: #ab639e;
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
    animation: fadeout 5s forwards;
}

@keyframes fadeout {
    to {
        opacity: 0;
        visibility: hidden;
    }
}
</style>