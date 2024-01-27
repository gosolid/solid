import { Transform } from 'stream';
import { expect } from 'chai';

describe('Transform Stream Tests', () => {
  it('should transform data correctly', (done) => {
    const transformStream = new Transform({
      transform(chunk, encoding, callback) {
        // Implement transformation logic
        const transformedData = chunk.toString().toUpperCase();
        this.push(transformedData);
        callback();
      },
    });

    let receivedData = '';

    transformStream.on('data', (chunk) => {
      // Handle transformed data
      receivedData += chunk.toString();
    });

    transformStream.on('end', () => {
      // Ensure data is transformed correctly
      expect(receivedData).to.equal('HELLO, WORLD!');
      done();
    });

    // Write data to the transform stream
    transformStream.write('Hello, ');
    transformStream.write('world!');
    transformStream.end();
  });

  it('should handle backpressure', (done) => {
    const transformStream = new Transform({
      transform(chunk, encoding, callback) {
        // Simulate a slow transformation process
        setTimeout(() => {
          const transformedData = chunk.toString().toUpperCase();
          this.push(transformedData);
          callback();
        }, 5);
      },
    });

    let receivedData = '';

    transformStream.on('data', (chunk) => {
      // Handle transformed data
      receivedData += chunk.toString();
    });

    transformStream.on('end', () => {
      // Ensure data is transformed correctly
      expect(receivedData).to.equal('x'.toUpperCase().repeat(100*100));
      done();
    });

    const largeData = 'x'.repeat(100);
    let i = 0;
    const next = () => {
        for (; i < 100; i++) {
          if(!transformStream.write(largeData)) {
            i++;
            return;
          }
        }

        transformStream.end();
    };

    transformStream.on('drain', () => {
        next();
    });

    next();
  }).timeout(5000); // Adjust the timeout based on the expected completion time

  it('should handle errors during transformation', (done) => {
    const transformStream = new Transform({
      transform(chunk, encoding, callback) {
        // Simulate an error during transformation
        callback(new Error('Transformation Error'));
      },
    });

    transformStream.on('error', (error) => {
      // Ensure error is handled correctly
      expect(error.message).to.equal('Transformation Error');
      done();
    });

    // Write data to the transform stream
    transformStream.write('Hello, world!');
    transformStream.end();
  });

  it('should handle object mode', (done) => {
    const transformStream = new Transform({
      objectMode: true,
      transform(chunk, encoding, callback) {
        // Implement object mode transformation logic
        const transformedData = { message: chunk.message.toUpperCase() };
        this.push(transformedData);
        callback();
      },
    });

    let receivedData: any[] = [];

    transformStream.on('data', (chunk) => {
      // Handle transformed data
      receivedData.push(chunk);
    });

    transformStream.on('end', () => {
      // Ensure object mode transformation is correct
      expect(receivedData).to.deep.equal([{ message: 'HELLO' }, { message: 'WORLD' }]);
      done();
    });

    // Write objects to the transform stream
    transformStream.write({ message: 'Hello' });
    transformStream.write({ message: 'World' });
    transformStream.end();
  });

  it('should handle multiple transformations', (done) => {
    const transformStream = new Transform({
      transform(chunk, encoding, callback) {
        // Implement transformation logic
        const transformedData = chunk.toString().toUpperCase();
        this.push(transformedData);
        callback();
      },
    });

    let receivedData = '';

    transformStream.on('data', (chunk) => {
      // Handle transformed data
      receivedData += chunk.toString();
    });

    transformStream.on('end', () => {
      // Ensure data is transformed correctly
      expect(receivedData).to.equal('HELLO, WORLD!HELLO, WORLD!');
      done();
    });

    // Write data to the transform stream multiple times
    transformStream.write('Hello, ');
    transformStream.write('world!');
    transformStream.write('Hello, ');
    transformStream.write('world!');
    transformStream.end();
  });

  it('should handle backpressure with object mode', (done) => {
    const transformStream = new Transform({
      objectMode: true,
      transform(chunk, encoding, callback) {
        // Simulate a slow transformation process
        setTimeout(() => {
          const transformedData = { message: chunk.message.toUpperCase() };
          this.push(transformedData);
          callback();
        }, 5);
      },
    });

    let receivedData: any[] = [];

    transformStream.on('data', (chunk) => {
      // Handle transformed data
      receivedData.push(chunk);
    });

    transformStream.on('end', () => {
      // Ensure object mode transformation is correct
      expect(receivedData).to.deep.equal(Array.from({ length: 100 }, () => ({ message: 'x'.toUpperCase().repeat(100) })));
      done();
    });

    // Write objects to the transform stream to trigger backpressure
    const largeData = { message: 'x'.repeat(100) };
    let i = 0;
    const next = () => {
        for (; i < 100; i++) {
          if(!transformStream.write(largeData)) {
            i++;
            return;
          }
        }

        transformStream.end();
    };

    transformStream.on('drain', () => {
        next();
    });

    next();
  });

  it('should handle errors during transformation in object mode', (done) => {
    const transformStream = new Transform({
      objectMode: true,
      transform(chunk, encoding, callback) {
        // Simulate an error during transformation
        callback(new Error('Transformation Error'));
      },
    });

    transformStream.on('error', (error) => {
      // Ensure error is handled correctly
      expect(error.message).to.equal('Transformation Error');
      done();
    });

    // Write objects to the transform stream
    transformStream.write({ message: 'Hello' });
    transformStream.write({ message: 'World' });
    transformStream.end();
  });

  it('should handle empty input data', (done) => {
    const transformStream = new Transform({
      transform(chunk, encoding, callback) {
        // Handle empty input data
        if (chunk.length === 0) {
          return callback();
        }

        // Implement transformation logic
        const transformedData = chunk.toString().toUpperCase();
        this.push(transformedData);
        callback();
      },
    });

    let receivedData = '';

    transformStream.on('data', (chunk) => {
      // Handle transformed data
      receivedData += chunk.toString();
    });

    transformStream.on('end', () => {
      // Ensure data is transformed correctly
      expect(receivedData).to.equal('HELLO, WORLD!');
      done();
    });

    // Write empty input data to the transform stream
    transformStream.write('');
    transformStream.write('Hello, ');
    transformStream.write('');
    transformStream.write('world!');
    transformStream.end();
  });
});
