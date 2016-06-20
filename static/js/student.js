/*var enrollApp angular.module("enroll", [])

	.controller("StudentController", function ($scope) {*/
	/*var self = this;
	$http.get('http://localhost:8080/v1/students').then(function(response)){
		self.students = response.data;
	}*/
var enrollApp = angular
						.module("enroll", [])
						.controller("StudentController", function ($scope, $http) {

							listStudent();
							function listStudent() {
								$http.get('http://localhost:8080/v1/students').then(function(response){
	  							$scope.students = response.data;
	  							});
							};



							$scope.addStudent = function () {
								var data = JSON.stringify({"name": $scope.studentInput}); 
								$http.post('http://localhost:8080/v1/students', data)
								.success(function(data){
									listStudent();
									$scope.studentInput = "";
								});
							};

							$scope.deleteStudent = function (uid) {
								if(confirm("Are you sure to delete ID "+uid+" ?")){
									//var uid = $scope.student.Id;
									var url = "http://localhost:8080/v1/students/"+uid
									$http.delete(url)
									.success(function(data){
										listStudent();
									});
								}
							};

							$scope.toggleStatus = function(uid, student) {
								//if(status=='2'){status='0';}else{status='2';}
								if(student.Status==2){student.Status=0;}else{student.Status=2;}
								var url = "http://localhost:8080/v1/students/"+uid
								var data = JSON.stringify(student)
								//Id: uid, Status: Status});
								$http.put(url, data)
								.success(function(data){
									listStudent();
								});
							};

						});
	
