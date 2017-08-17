angular.module('myapp', [])
	.run(function() {
		console.log('App running');
	})
	// DIRECTICA QUE CONVIERTE A MAYUSCULAS EL TEXTO EN UN COMPONENTE
	.directive('capitalize', function () {
		return {
			require: 'ngModel',
			link: function (scope, element, attrs, modelCtrl) {
				var capitalize = function (inputValue) {
					if (inputValue == undefined) inputValue = '';
					var capitalized = inputValue.toUpperCase();
					if (capitalized !== inputValue) {
						modelCtrl.$setViewValue(capitalized);
						modelCtrl.$render();
					}
					return capitalized;
				}
				modelCtrl.$parsers.push(capitalize);
				capitalize(scope[attrs.ngModel]);
			}
		};
	})
	// DIRECTICA QUE CONVIERTE A MINUSCULAS EL TEXTO EN UN COMPONENTE
	.directive('uncapitalize', function () {
		return {
			require: 'ngModel',
			link: function (scope, element, attrs, modelCtrl) {
				var nocapitalize = function (inputValue) {
					if (inputValue == undefined) inputValue = '';
					var nocapitalized = inputValue.toLowerCase();
					if (nocapitalized !== inputValue) {
						modelCtrl.$setViewValue(nocapitalized);
						modelCtrl.$render();
					}
					return nocapitalized;
				}
				modelCtrl.$parsers.push(nocapitalize);
				nocapitalize(scope[attrs.ngModel]);
			}
		};
	})
	// DIRECTIVA QUE VALIDA QUE SOLO SE INGRESEN NÚMEROS
	.directive('validNumber', function () {
		return {
			require: '?ngModel',
			link: function (scope, element, attrs, ngModelCtrl) {
				if (!ngModelCtrl) {
					return;
				}

				ngModelCtrl.$parsers.push(function (val) {
					if (angular.isUndefined(val)) {
						var val = '';
					}
					var clean = val.replace(/[^0-9]+/g, '');
					if (val !== clean) {
						ngModelCtrl.$setViewValue(clean);
						ngModelCtrl.$render();
					}
					return clean;
				});

				element.bind('keypress', function (event) {
					if (event.keyCode === 32) {
						event.preventDefault();
					}
				});
			}
		};
	});