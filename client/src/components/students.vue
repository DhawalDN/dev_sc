<template>
	<div class="container">
		<div class="alert alert-success" role="alert" v-if="isSubmitted">
			<h4 class="alert-heading">Well done!</h4>
			<p>You have Updated the User</p>
			
			
		</div>
		<div class="container d-flex w100">
			<div class="wrap-table100">
				<div class="table100 ver1 m-b-110">
					<div class="table100-head">
						<table>
							<thead>
								<tr class="row100 head">
									<th class="cell100 column1">Name</th>
									<th class="cell100 column2">Email</th>
									<th class="cell100 column3">Gender</th>
									<th class="cell100 column4"></th>
									
								</tr>
							</thead>
						</table>
					</div>

					<div class="table100-body js-pscroll">
						<table>
							<tbody>
								<tr class="row100 body" v-for="(result,index) in results.data" :key="index">
									<td class="cell100 column1">{{result.name}}</td>
									<td class="cell100 column2">{{result.email}}</td>
									<td class="cell100 column3">{{result.gender}}</td>
									<td class="cell100 column4">
										<button
											type="button"
											class="btn btn-success"
											data-toggle="modal"
											data-target="#exampleModal"
											@click="student_Data=result"
										>Edit</button>&nbsp;
										
									
										<button
											type="button"
											class="btn btn-danger"
											@click="Delete(result.id)"
										>Delete</button>
									</td>
								</tr>
							</tbody>
						</table>
					</div>
				</div>
			</div>
		</div>
		
		<div
			class="modal fade"
			id="exampleModal"
			tabindex="-1"
			role="dialog"
			aria-labelledby="exampleModalLabel"
			aria-hidden="true"
		>
			<div class="modal-dialog" role="document">
				<div class="modal-content">
					<div class="modal-header">
						<h5 class="modal-title" id="exampleModalLabel" >Update Form</h5>
						<button type="button" class="close" data-dismiss="modal" aria-label="Close">
							<span aria-hidden="true">&times;</span>
						</button>
					</div>
					<div class="modal-body">
						<form class="card-body">
							<div class="form-group">
								<div class="form-group">
									<label for="fname">Full Name</label>
									<input type="text" id="email" class="form-control border" v-model="student_Data.name" />
								</div>
								<div class="form-group">
									<label for="email">E-mail</label>
									<input type="text" id="email" class="form-control border" v-model="student_Data.email" />
								</div>
								<div class="form-group">
									<div class>
										<div class="form-check form-check-inline">
											<input
												class="form-check-input"
												type="radio"
												name="inlineRadioOptions"
												id="inlineRadio1"
												value="option1"
											/>
											<label class="form-check-label" for="inlineRadio1">Male</label>
										</div>
										<div class="form-check form-check-inline">
											<input
												class="form-check-input"
												type="radio"
												name="inlineRadioOptions"
												id="inlineRadio2"
												value="option2"
											/>
											<label class="form-check-label" for="inlineRadio2">Female</label>
										</div>
									</div>
								</div>
							</div>
						</form>
					</div>
					<div class="modal-footer">
						<button type="button" class="btn btn-primary" data-dismiss="modal" @click="Update()">Save changes</button>
						<button type="button" class="btn btn-secondary" data-dismiss="modal">Close</button>
					</div>
				</div>
			</div>
		</div>
		
	</div>
</template>

<script>
	export default {
		name: "StudentTable",
		data() {
			return {
				results: [],
				u_id: {
					id:""
				},
				// student_Data: {
				// 	id:"",
				// 	email: "",
				// 	name: "",
				// 	gender: "Male"
				// },
				student_Data:{},
				isSubmitted: false
			};
		},
		methods: {
			Update() {
				console.log(this.student_Data);
				
				delete this.student_Data.password
				this.axios.post("/api/v1/update", this.student_Data).then(response => {
					console.log(response.data);
					this.isSubmitted=true
					
				});
			},
			
			Delete(uid) {
				this.u_id.id=uid
				console.log(this.u_id)
				this.axios.post("/api/v1/delete", this.u_id).then(response => {
					console.log(response.data);
					this.Fetch();
					
				})
			},
			Fetch() {
				this.axios.get("/api/v1/fetch").then(response => {
				//console.log(response.data);
				this.results = response.data;

				console.log("fetched")
				// console.log(this.results);
			});
			}
		},

		mounted() {
			console.log("mounted");
			this.Fetch()
		},
		created() {
			console.log("created");
			// this.Fetch()
		}
	};
</script>

