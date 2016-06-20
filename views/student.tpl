<!DOCTYPE html>

<html ng-app="enroll">
    <head>
        <title>Student Registration</title>
        <meta http-equiv="Content-Type" content="text/html; charset=utf-8">
        <script src="/static/js/angular.min.js"></script>
        <script src="/static/js/student.js"></script>
    </head>
    <body ng-controller="StudentController">
        <header class="hero-unit" style="background-color:#A9F16C">
            <div class="container">
            <div class="row">
              <div class="hero-text">
                <h1>Welcome to DMCE!</h1>
                <ul>
    				<li ng-repeat="student in students">
      					<span>{{.name}}</span>
      				</li>
  				</ul>
              </div>
            </div>
            </div>
        </header>
    </body>
</html>