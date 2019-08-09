<template>
	<div class="card border-0">
		<div class="card-body">
			<form class="card-body">
				<div class="row">
					<div class="form-group">
						<h1>Student Registration</h1>
						<hr />
						<div class="form-group">
							<label for="fname">Full Name</label>
							<input type="text" id="email" class="form-control border" v-model="studentData.name" />
						</div>
						<div class="form-group">
							<label for="email">E-mail</label>
							<input type="text" id="email" class="form-control border" v-model="studentData.email" />
						</div>
						<div class="form-group">
							<label for="password">Password</label>
							<input type="password" id="password" class="form-control border" v-model="studentData.password" />
						</div>
						<div class="form-group">
							<label for="age">Age</label>
							<input type="number" id="age" class="form-control border" v-model="studentData.age" />
						</div>
					</div>
				</div>

				<div class="row ">
					<div class="form-group border-0">
						
						<label for="male">
							<input type="radio" id="male" value="Male" class="form-control" v-model="studentData.gender" /> Male
						</label>
						<label for="female">
							<input type="radio" id="female" value="Female" class="form-control" v-model="studentData.gender" /> Female
						</label>
					</div>
				</div>

				<hr />
				<div class="row">
					<div class="col-xs-12 col-sm-8 col-sm-offset-2 col-md-6 col-md-offset-3">
						<button class="btn btn-primary" @click.prevent="submitted">Submit!</button>
					</div>
				</div>
			</form>
		</div>
		<hr />
		<div class="row" v-if="isSubmitted">
			<div class="col-xs-12 col-sm-8 col-sm-offset-2 col-md-6 col-md-offset-3">
				<div class="panel panel-default">
					<div class="panel-heading">
						<h4>Student Data</h4>
					</div>
					<div class="panel-body" v-for="item in row" :key="item">
						<p>Full Name: {{ item.name }}</p>
						<p>Email:{{ item.email }}</p>
						<p>Age: {{ item.age }}</p>
						<p>Gender:{{ item.gender}}</p>
					</div>
				</div>
			</div>
		</div>
	</div>
</template>

<script>
	export default {
		name: "register",
		data() {
			return {
				studentData: {
					email: "",
					name: "",
					password: "",
					age: 25,
					gender: "Male"
				},
				row: [],
				isSubmitted: false
			};
		},
		methods: {
			submitted() {
				this.isSubmitted = true;
				
				this.axios.post("/api/v1/create", this.studentData).then(res => {
                    this.row.push(this.studentData)
                    this.console.log(res)
				});
			}
        },
        
	};
</script>
