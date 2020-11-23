const gulp = require('gulp');
const notify = require("gulp-notify");
const plumber = require("gulp-plumber");
const pug = require('gulp-pug');
const sass = require('gulp-sass');
const typescript = require('gulp-typescript');

const paths = {
    'pug': './src/pug/',
    'sass': './src/scss/',
    'typescript': './src/ts/',
    'html': './dist/',
    'css': './dist/',
    'javascript': './dist/'
};

// Pug → HTML
const pugCompile = () => {
    return gulp.src([
        paths.pug + '**/*.pug'
    ])
    .pipe(
        plumber({
            errorHandler: notify.onError("Error: <%= error.message %>")
        })
    )
    .pipe(
        pug({pretty: false})
    )
    .pipe(
        gulp.dest(paths.html)
    )
};

// Sass(SCSS) → CSS
const sassCompile = () => {
    return gulp.src([
        paths.sass + '**/*.scss'
    ])
    .pipe(
        plumber({
            errorHandler: notify.onError("Error: <%= error.message %>")
        })
    )
    .pipe(
        sass({
            outputStyle: 'compressed'
        })
    )
    .pipe(
        gulp.dest(paths.css)
    )
};

// TypeScript → JavaScript
const tsCompile = () => {
    return gulp.src([
        paths.typescript + '**/*.ts',
        '!./node_modules/**'
    ])
    .pipe(
        plumber({
            errorHandler: notify.onError("Error: <%= error.message %>")
        })
    )
    .pipe(
        typescript({
            module: 'commonjs',
            target: 'ES5'
        })
    )
    .pipe(
        gulp.dest(paths.javascript)
    )
};

// 自動コンパイル
const liveCompile = () => {
    gulp.watch(paths.pug + '**/*.pug')
    .on('change', gulp.series(pugCompile));

    gulp.watch(paths.sass + '**/*.scss')
    .on('change', gulp.series(sassCompile));

    gulp.watch(paths.typescript + '**/*.ts')
    .on('change', gulp.series(tsCompile));
}

gulp.task('default', gulp.parallel(pugCompile, sassCompile, tsCompile));
gulp.task('live', liveCompile);