---
title: React-Native-Doc-Animated
date: 2018-07-31 15:14:50
tags:
---

The `Animated` library is designed to make animations fluid, powerful, and easy to build and maintain. `Animated` focuses on declarative relationships between inputs and outputs, with configurable transforms in between, and simple `start`/`stop` methods to control time-based animation execution.

Animated动画库旨在让动画流畅,强大并且易于构建和使用. 
Animated注重在动画过程的输入和输出之间使用变换设置, 并且使用简单的开始/结束方法控制基于时间的动画执行过程.

The simplest workflow for creating an animation is to create an `Animated.Value`, hook it up to one or more style attributes of an animated component, and then drive updates via animations using `Animated.timing()`:

创建一个动画最简单的方法: 
1. 创建一个动画值.`Animated.Value(0)`;
2. 在可动画组件的一个或多个属性上引用这个值.
3. 通过`Animated.timing()`函数动态更新动画值.

```javascript
Animated.timing(
  // Animate value over time
  this.state.fadeAnim, // The value to drive
  {
    toValue: 1, // Animate to final value of 1
  }
).start(); // Start the animation
```

Refer to the [Animations](animations.md#animated-api) guide to see additional examples of `Animated` in action.

## Overview 概览

There are two value types you can use with `Animated`:

* [`Animated.Value()`](animated.md#value) for single values
* [`Animated.ValueXY()`](animated.md#valuexy) for vectors

`Animated`有两种动画值: 
1. `Animated.Value()` 用于单个值;
2. `Animated.ValueXY()` 用于向量值;

`Animated.Value` can bind to style properties or other props, and can be interpolated as well. A single `Animated.Value` can drive any number of properties.

`Animated.Value()` 可以绑定到组件`style`属性或者`props`上, 也可以使用插值器.

一个单值`Animated.Value()`可以驱动多个属性值.

### Configuring animations  配置动画

`Animated` provides three types of animation types. Each animation type provides a particular animation curve that controls how your values animate from their initial value to the final value:

* [`Animated.decay()`](animated.md#decay) starts with an initial velocity and gradually slows to a stop.
* [`Animated.spring()`](animated.md#spring) provides a simple spring physics model.
* [`Animated.timing()`](animated.md#timing) animates a value over time using [easing functions](easing.md).

In most cases, you will be using `timing()`. By default, it uses a symmetric easeInOut curve that conveys the gradual acceleration of an object to full speed and concludes by gradually decelerating to a stop.

`Animated` 提供了三种类型的动画. 每种类型提供了一个特定的动画方式,以便能让您使用初始值和最终值控制动画的执行.

1. `Animated.dacay()` 减速动画, 从一个初始速度开始然后渐渐地减速.
2. `Animated.spring()` 回弹动画, 提供一个简单的物理回弹模型.
3. `Animated.timing()` 时间动画, 随时间变化执行动画.

大多是情况下,我们只是用`timing()`动画. 
默认情况下，它使用对称的EaseNoUT曲线，将物体的逐渐加速传递到全速，并通过逐渐减速到停止而结束。

### Working with animations 让动画开始工作

Animations are started by calling `start()` on your animation. `start()` takes a completion callback that will be called when the animation is done. If the animation finished running normally, the completion callback will be invoked with `{finished: true}`. If the animation is done because `stop()` was called on it before it could finish (e.g. because it was interrupted by a gesture or another animation), then it will receive `{finished: false}`.

开始和结束方法都有回调,回调字段有个finish标记是否已经完成动画.

猜想是下面的逻辑.
```
let testAnimated = new Animated.Value(0)
Animated.timing(this.state.testAnimated,
{
    toValue:1,
}).start((callback) => {
    if(callback.finished) {
        alert('动画执行完毕')
    }
})
```

### Using the native driver 使用原生驱动

By using the native driver, we send everything about the animation to native before starting the animation, allowing native code to perform the animation on the UI thread without having to go through the bridge on every frame. Once the animation has started, the JS thread can be blocked without affecting the animation.

You can use the native driver by specifying `useNativeDriver: true` in your animation configuration. See the [Animations](animations.md#using-the-native-driver) guide to learn more.

### Animatable components 可动画组件

Only animatable components can be animated. These special components do the magic of binding the animated values to the properties, and do targeted native updates to avoid the cost of the react render and reconciliation process on every frame. They also handle cleanup on unmount so they are safe by default.

* [`createAnimatedComponent()`](animated.md#createanimatedcomponent) can be used to make a component animatable.

`Animated` exports the following animatable components using the above wrapper:

* `Animated.Image`
* `Animated.ScrollView`
* `Animated.Text`
* `Animated.View`

### Composing animations 组合动画

Animations can also be combined in complex ways using composition functions:

* [`Animated.delay()`](animated.md#delay) starts an animation after a given delay.
一定延时后启动动画
* [`Animated.parallel()`](animated.md#parallel) starts a number of animations at the same time.
同时启动多个动画
* [`Animated.sequence()`](animated.md#sequence) starts the animations in order, waiting for each to complete before starting the next.
队列启动动画.
* [`Animated.stagger()`](animated.md#stagger) starts animations in order and in parallel, but with successive delays.
按策略启动动画.

Animations can also be chained together simply by setting the `toValue` of one animation to be another `Animated.Value`. See [Tracking dynamic values](animations.md#tracking-dynamic-values) in the Animations guide.
也可以简单是使用toValue使得多个动画衔接执行.

By default, if one animation is stopped or interrupted, then all other animations in the group are also stopped.

默认情况下, 如果某一个动画停止了或者被拦截器拦截了 其他的同组动画都会停止.


### Combining animated values  组合值动画

You can combine two animated values via addition, subtraction, multiplication, division, or modulo to make a new animated value:

* [`Animated.add()`](animated.md#add)
* [`Animated.subtract()`](animated.md#subtract)
* [`Animated.divide()`](animated.md#divide)
* [`Animated.modulo()`](animated.md#modulo)
* [`Animated.multiply()`](animated.md#multiply)

可以使用上面的四个加减乘除的方法产生新的动画值.

### Interpolation  插值器.

The `interpolate()` function allows input ranges to map to different output ranges. By default, it will extrapolate the curve beyond the ranges given, but you can also have it clamp the output value. It uses lineal interpolation by default but also supports easing functions.

* [`interpolate()`](animated.md#interpolate)

Read more about interpolation in the [Animation](animations.md#interpolation) guide.

`interpolate()`函数允许输入范围映射到不同的输出范围。默认情况下，它会将曲线外推到给定的范围之外，但也可以将输出值箝位。它默认使用线性插值，但也支持`Ease`函数。

### Handling gestures and other events 处理手势和事件

Gestures, like panning or scrolling, and other events can map directly to animated values using `Animated.event()`. This is done with a structured map syntax so that values can be extracted from complex event objects. The first level is an array to allow mapping across multiple args, and that array contains nested objects.

* [`Animated.event()`](animated.md#event)

For example, when working with horizontal scrolling gestures, you would do the following in order to map `event.nativeEvent.contentOffset.x` to `scrollX` (an `Animated.Value`):

```javascript
 onScroll={Animated.event(
   // scrollX = e.nativeEvent.contentOffset.x
   [{ nativeEvent: {
        contentOffset: {
          x: scrollX
        }
      }
    }]
 )}
```

### Methods

* [`decay`](animated.md#decay)
* [`timing`](animated.md#timing)
* [`spring`](animated.md#spring)
* [`add`](animated.md#add)
* [`subtract`](animated.md#subtract)
* [`divide`](animated.md#divide)
* [`multiply`](animated.md#multiply)
* [`modulo`](animated.md#modulo)
* [`diffClamp`](animated.md#diffclamp)
* [`delay`](animated.md#delay)
* [`sequence`](animated.md#sequence)
* [`parallel`](animated.md#parallel)
* [`stagger`](animated.md#stagger)
* [`loop`](animated.md#loop)
* [`event`](animated.md#event)
* [`forkEvent`](animated.md#forkevent)
* [`unforkEvent`](animated.md#unforkevent)

### Properties

* [`Value`](animated.md#value)
* [`ValueXY`](animated.md#valuexy)
* [`Interpolation`](animated.md#interpolation)
* [`Node`](animated.md#node)
* [`createAnimatedComponent`](animated.md#createanimatedcomponent)
* [`attachNativeEvent`](animated.md#attachnativeevent)

---

# Reference

## Methods

When the given value is a ValueXY instead of a Value, each config option may be a vector of the form `{x: ..., y: ...}` instead of a scalar.

### `decay()`

```javascript
static decay(value, config)
```

Animates a value from an initial velocity to zero based on a decay coefficient.

Config is an object that may have the following options:

* `velocity`: Initial velocity. Required.
* `deceleration`: Rate of decay. Default 0.997.
* `isInteraction`: Whether or not this animation creates an "interaction handle" on the `InteractionManager`. Default true.
* `useNativeDriver`: Uses the native driver when true. Default false.

---

### `timing()`

```javascript
static timing(value, config)
```

Animates a value along a timed easing curve. The [`Easing`](easing.md) module has tons of predefined curves, or you can use your own function.

Config is an object that may have the following options:

* `duration`: Length of animation (milliseconds). Default 500.
* `easing`: Easing function to define curve. Default is `Easing.inOut(Easing.ease)`.
* `delay`: Start the animation after delay (milliseconds). Default 0.
* `isInteraction`: Whether or not this animation creates an "interaction handle" on the `InteractionManager`. Default true.
* `useNativeDriver`: Uses the native driver when true. Default false.

---

### `spring()`

```javascript
static spring(value, config)
```

Animates a value according to an analytical spring model based on [damped harmonic oscillation](https://en.wikipedia.org/wiki/Harmonic_oscillator#Damped_harmonic_oscillator). Tracks velocity state to create fluid motions as the `toValue` updates, and can be chained together.

Config is an object that may have the following options.

Note that you can only define one of bounciness/speed, tension/friction, or stiffness/damping/mass, but not more than one:

The friction/tension or bounciness/speed options match the spring model in [Facebook Pop](https://github.com/facebook/pop), [Rebound](http://facebook.github.io/rebound/), and [Origami](http://origami.design/).

* `friction`: Controls "bounciness"/overshoot. Default 7.
* `tension`: Controls speed. Default 40.
* `speed`: Controls speed of the animation. Default 12.
* `bounciness`: Controls bounciness. Default 8.

Specifying stiffness/damping/mass as parameters makes `Animated.spring` use an analytical spring model based on the motion equations of a [damped harmonic oscillator](https://en.wikipedia.org/wiki/Harmonic_oscillator#Damped_harmonic_oscillator). This behavior is slightly more precise and faithful to the physics behind spring dynamics, and closely mimics the implementation in iOS's CASpringAnimation primitive.

* `stiffness`: The spring stiffness coefficient. Default 100.
* `damping`: Defines how the spring’s motion should be damped due to the forces of friction. Default 10.
* `mass`: The mass of the object attached to the end of the spring. Default 1.

Other configuration options are as follows:

* `velocity`: The initial velocity of the object attached to the spring. Default 0 (object is at rest).
* `overshootClamping`: Boolean indiciating whether the spring should be clamped and not bounce. Default false.
* `restDisplacementThreshold`: The threshold of displacement from rest below which the spring should be considered at rest. Default 0.001.
* `restSpeedThreshold`: The speed at which the spring should be considered at rest in pixels per second. Default 0.001.
* `delay`: Start the animation after delay (milliseconds). Default 0.
* `isInteraction`: Whether or not this animation creates an "interaction handle" on the `InteractionManager`. Default true.
* `useNativeDriver`: Uses the native driver when true. Default false.

---

### `add()`

```javascript
static add(a, b)
```

Creates a new Animated value composed from two Animated values added together.

---

### `subtract()`

```javascript
static subtract(a, b)
```

Creates a new Animated value composed by subtracting the second Animated value from the first Animated value.

---

### `divide()`

```javascript
static divide(a, b)
```

Creates a new Animated value composed by dividing the first Animated value by the second Animated value.

---

### `multiply()`

```javascript
static multiply(a, b)
```

Creates a new Animated value composed from two Animated values multiplied together.

---

### `modulo()`

```javascript
static modulo(a, modulus)
```

Creates a new Animated value that is the (non-negative) modulo of the provided Animated value

---

### `diffClamp()`

```javascript
static diffClamp(a, min, max)
```

Create a new Animated value that is limited between 2 values. It uses the difference between the last value so even if the value is far from the bounds it will start changing when the value starts getting closer again. (`value = clamp(value + diff, min, max)`).

This is useful with scroll events, for example, to show the navbar when scrolling up and to hide it when scrolling down.

---

### `delay()`

```javascript
static delay(time)
```

Starts an animation after the given delay.

---

### `sequence()`

```javascript
static sequence(animations)
```

Starts an array of animations in order, waiting for each to complete before starting the next. If the current running animation is stopped, no following animations will be started.

---

### `parallel()`

```javascript
static parallel(animations, config?)
```

Starts an array of animations all at the same time. By default, if one of the animations is stopped, they will all be stopped. You can override this with the `stopTogether` flag.

---

### `stagger()`

```javascript
static stagger(time, animations)
```

Array of animations may run in parallel (overlap), but are started in sequence with successive delays. Nice for doing trailing effects.

---

### `loop()`

```javascript
static loop(animation)
```

Loops a given animation continuously, so that each time it reaches the end, it resets and begins again from the start. Can specify number of times to loop using the key `iterations` in the config. Will loop without blocking the UI thread if the child animation is set to `useNativeDriver: true`. In addition, loops can prevent `VirtualizedList`-based components from rendering more rows while the animation is running. You can pass `isInteraction: false` in the child animation config to fix this.

---

### `event()`

```javascript
static event(argMapping, config?)
```

Takes an array of mappings and extracts values from each arg accordingly, then calls `setValue` on the mapped outputs. e.g.

```javascript
 onScroll={Animated.event(
   [{nativeEvent: {contentOffset: {x: this._scrollX}}}],
   {listener: (event) => console.log(event)}, // Optional async listener
 )}
 ...
 onPanResponderMove: Animated.event([
   null,                // raw event arg ignored
   {dx: this._panX}],    // gestureState arg
{listener: (event, gestureState) => console.log(event, gestureState)}, // Optional async listener
 ),
```

Config is an object that may have the following options:

* `listener`: Optional async listener.
* `useNativeDriver`: Uses the native driver when true. Default false.

---

### `forkEvent()`

```javascript
static forkEvent(event, listener)
```

Advanced imperative API for snooping on animated events that are passed in through props. It permits to add a new javascript listener to an existing `AnimatedEvent`. If `animatedEvent` is a simple javascript listener, it will merge the 2 listeners into a single one, and if `animatedEvent` is null/undefined, it will assign the javascript listener directly. Use values directly where possible.

---

### `unforkEvent()`

```javascript
static unforkEvent(event, listener)
```

## Properties

---

---

---

---

---