<template>
	<div class="container">
		<nav class="navbar navbar-expand-lg navbar-light bg-light">
			<a class="navbar-brand" href="#">SC</a>
			<button
				class="navbar-toggler"
				type="button"
				data-toggle="collapse"
				data-target="#navbarNavAltMarkup"
				aria-controls="navbarNavAltMarkup"
				aria-expanded="false"
				aria-label="Toggle navigation"
			>
				<span class="navbar-toggler-icon"></span>
			</button>
			<div class="collapse navbar-collapse" id="navbarNavAltMarkup">
				<div class="navbar-nav">
					<a class="nav-item nav-link active" href="#">
						<router-link to="/" >Register</router-link>
						<span class="sr-only">(current)</span>
					</a>
					<a class="nav-item nav-link" href="#">
						
                        <router-link to="/students" exact>Student Table</router-link>
					</a>
					
				</div>
			</div>
		</nav>

		<router-view />
	</div>
</template>

<script>
	export default {
		name: "App"
	};
</script>

<style>
</style>