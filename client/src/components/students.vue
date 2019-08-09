<template>
	<div class="container">
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
								</tr>
							</thead>
						</table>
					</div>
					

					<div class="table100-body js-pscroll" >
						<table>
							<tbody>
								<tr class="row100 body" v-for="result in results.data" :key="result">
									<td class="cell100 column1">{{result.name}}</td>
									<td class="cell100 column2">{{result.email}}</td>
									<td class="cell100 column3">{{result.gender}}</td>
								</tr>
							</tbody>
						</table>
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
				results:[]
			};
		},
		mounted() {
			// console.log("sfadsf");
			this.axios.get("/api/v1/fetch").then(response => {
				//console.log(response.data);
				this.results = response.data;
				// console.log(this.results);
			});
		}
	};
</script>

